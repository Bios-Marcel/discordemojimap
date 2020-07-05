package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// USE FILE:
// https://discord.com/assets/b38205c8085075585265.js

const goCode = `package discordemojimap

var EmojiMap = map[string]string {
%s}
`

// emojiJSONRegex matches the emoji JSON in a certain asset file. This JSON can
// have one of its first object key as one of the strings matched in this OR'd
// regex.
var emojiJSONRegex = regexp.MustCompile(`'{"(people|activity|flags|food|nature|objects|symbols|travel)":.*}'`)

type EmojiGroups map[string][]Emoji

func (eg EmojiGroups) GroupNames() []string {
	var keys = make([]string, 0, len(eg))
	for key := range eg {
		keys = append(keys, key)
	}
	return keys
}

type Emoji struct {
	Names          []string `json:"names"`
	Surrogates     string   `json:"surrogates"` // the emoji
	UnicodeVersion float64  `json:"unicodeVersion"`

	// Optionals

	HasDiversity      bool    `json:"hasDiversity,omitempty"`
	HasMultiDiversity bool    `json:"hasMultiDiversity,omitempty"`
	Diversities       []Emoji `json:"diversityChildren,omitempty"`
}

// GoSyntax writes the representation of the emoji as a single map entry.
func (e Emoji) GoSyntax(w io.Writer) (n int, err error) {
	for _, emojiName := range e.Names {
		// Use %+q so Unicode emojis are formatted as \U or \u.
		w, err := fmt.Fprintf(w, "\t%q: %+q,\n", emojiName, e.Surrogates)
		n += w // accumulate bytes written
		if err != nil {
			return n, err
		}
	}

	return n, nil
}

func main() {
	var path = ""
	flag.StringVar(&path, "path", "", "path should be a relative or absolute path to the file to create the mapping from.")
	var out = "mapping.go"
	flag.StringVar(&out, "out", out, "output file path")

	flag.Parse()

	if path == "" {
		log.Fatalln("Usage:", filepath.Base(os.Args[0]), "-path <file>")
	}

	absPath, absPathError := filepath.Abs(path)
	if absPathError != nil {
		log.Fatalf("Error retrieving absolute filepath for input: %q: %s\n", path, absPathError)
	}

	data, readError := ioutil.ReadFile(absPath)
	if readError != nil {
		log.Fatalf("Error reading data: %s\n", readError)
	}

	emojiJSON := emojiJSONRegex.Find(data)
	if emojiJSON == nil {
		log.Fatalln("Emojis JSON not found.")
	}

	// Trim the single quotes matched.
	emojiJSON = bytes.Trim(emojiJSON, "'")

	var groups EmojiGroups

	if err := json.Unmarshal(emojiJSON, &groups); err != nil {
		log.Fatalln("Failed to unmarshal JSON:", err)
	}

	// Even though maps are unordered, we'd preferably still want a reproducible
	// output.
	var names = groups.GroupNames()
	sort.Strings(names)

	var mapping strings.Builder
	for _, name := range names {
		for _, emoji := range groups[name] {
			// Write the basic emojis.
			emoji.GoSyntax(&mapping)

			// Check if we have toned emojis. Write all of them if we do.
			for _, emoji := range emoji.Diversities {
				emoji.GoSyntax(&mapping)
			}
		}
	}

	f, err := os.OpenFile(out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Failed to open output:", err)
	}
	defer f.Close()

	if _, err := fmt.Fprintf(f, goCode, mapping.String()); err != nil {
		log.Fatalln("Failed to format Go code:", err)
	}
}

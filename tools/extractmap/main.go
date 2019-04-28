package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

const goCode = `package discordemojimap

var emojiMap = map[string]string {%s
}
`

var emojiPairRegex = regexp.MustCompile("{names:\\[(.*?)\\],surrogates:\"(.*?)\"")

func main() {
	var path string
	flag.StringVar(&path, "path", "", "path should be a relative or absolute path to the file to create the mapping from.")
	flag.Parse()

	if path == "" {
		log.Fatalln("Path must not be empty.")
	}

	absPath, absPathError := filepath.Abs(path)
	if absPathError != nil {
		log.Fatalf("Error retrieving absolute filepath for input: '%s': %s", path, absPathError)
	}

	data, readError := ioutil.ReadFile(absPath)
	if readError != nil {
		log.Fatalf("Error reading data: %s", readError)
	}

	dataAsString := string(data)
	matches := emojiPairRegex.FindAllStringSubmatch(dataAsString, -1)

	mapping := ""
	for _, match := range matches {
		value := match[2]
		keys := strings.Split(match[1], ",")
		for _, key := range keys {
			mapping = fmt.Sprintf("%s\n\t%s: \"%s\",", mapping, key, value)
		}
	}

	fileContent := fmt.Sprintf(goCode, mapping)

	ioutil.WriteFile("mapping.go", []byte(fileContent), 0755)
}

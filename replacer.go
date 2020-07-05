// Package discordemojimap provides a Replace function in order to escape
// emoji sequences with their respective emojis.
package discordemojimap

import (
	"regexp"
	"strings"
)

var emojiCodeRegex = regexp.MustCompile("(?s):[a-zA-Z0-9_]+:")

// Replace all emoji sequences contained in the discord emoji map with their
// respective emojis.
func Replace(input string) string {
	// Return the input as-is if it has less than a pair of colons.
	if len(input) <= 2 {
		return input
	}

	return emojiCodeRegex.ReplaceAllStringFunc(input, func(match string) string {
		emojified, contains := EmojiMap[strings.ToLower(match[1:len(match)-1])]
		if !contains {
			return match
		}

		return emojified
	})
}

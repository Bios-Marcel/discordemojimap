package discordemojimap

import "regexp"

var emojiCodeRegex = regexp.MustCompile("(?s):[a-zA-Z0-9_]+:")

// Replace all emoji sequences contained in the discord emoji map with their
// respective emojis.
//
// Examples for valid input:
//     Replace("Hello World :sun_with_face:")
// would result in
//     "Hello World 🌞"
func Replace(input string) string {
	if len(input) <= 2 {
		return input
	}

	replacedEmojis := emojiCodeRegex.ReplaceAllStringFunc(input, func(match string) string {
		emojified, contains := emojiMap[match[1:len(match)-1]]
		if !contains {
			return match
		}

		return emojified
	})

	return replacedEmojis
}

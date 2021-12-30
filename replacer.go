// Package discordemojimap provides a Replace function in order to escape
// emoji sequences with their respective emojis.
package discordemojimap

import "strings"

// Replace all emoji sequences contained in the discord emoji map with their
// respective emojis. For example:
//     fmt.Println(Replace("Hello World :sun_with_face:"))
//     //Output: Hello World 🌞
// This function is optimized for lowercased emoji sequence, meaning that
// sequences such as ":SUNGLASSES:" will consume slightly more memory and be
// slightly slower. However, the impact should be insignificant in most cases.
func Replace(input string) string {
	// Return the input as-is if it has less than a pair of colons.
	if len(input) <= 2 {
		return input
	}

	// In order to avoid wasting memory, we don't use bytes.buffer, but do it ourselves.
	var buffer []byte

	start := -1
	var lastEnd int
	// Instead of for ranging over the string, we treat it as a byte array to
	// save CPU cycles.
	for index := 0; index < len(input); index++ {
		char := input[index]
		// Even though we might get codepoints out of the ascii range, one
		// byte of a unicode codepoint can never be a colon. This is proven
		// by the test TestThatPartOfARuneCannotBeColon and probably wouldn't
		// need proving if I understood unicode and UTF-8 better.
		if char != ':' {
			continue
		}

		if start == -1 {
			start = index
			continue
		}

		// Occurence of something like "Hello :: World", in which case we needn't do anything.
		if index-start == 1 {
			start = -1
			continue
		}

		emojiSequence := input[start+1 : index]
		emojified, contains := EmojiMap[emojiSequence]
		if !contains {
			// Since the previous check is case sensitive, we do the same in a case
			// insensitive manner to make use of the best case performance.
			emojiSequence = strings.ToLower(input[start+1 : index])
			emojified, contains = EmojiMap[emojiSequence]
			if !contains {
				start = -1
				// Solves cases such as this ":sunglassesö:sunglasses:", where
				// the sequence wouldn't be sucessfully resolved otherwise.
				// Danke Marvin.
				index--
				continue
			}
		}

		if len(buffer) == 0 {
			// Potentially allocate a bit more than required, but not having to reallocate
			buffer = make([]byte, 0, len(input)-len(emojiSequence)-2+len(emojified))
		}
		buffer = append(buffer, input[lastEnd:start]...)
		buffer = append(buffer, emojified...)
		start = -1
		lastEnd = index + 1
	}

	// Since we only ever append after we've found a matching
	// sequence, we still have to append what's left over.
	if lastEnd > 0 {
		buffer = append(buffer, input[lastEnd:]...)
		return string(buffer)
	}

	// Since lastEnd is always set if we've had a match, we don't need to
	// check the buffer content anymore and can directly fallback to the input.
	return input
}

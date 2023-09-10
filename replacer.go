// Package discordemojimap provides a Replace function in order to escape
// emoji sequences with their respective emojis.
package discordemojimap

// Replace all emoji sequences contained in the discord emoji map with their
// respective emojis. For example:
//
//	fmt.Println(Replace("Hello World :sun_with_face:"))
//	//Output: Hello World 🌞
//
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

		// Start of new sequence or an occurence of something like "::+1:" in
		// which case we want to ignore the first colon and start again from
		// the next one, as it might be a valid sequence.
		if start == -1 || index-start == 1 {
			start = index
			continue
		}

		emojiSequence := input[start+1 : index]
		if lowered := toLower(emojiSequence); lowered != "" {
			emojiSequence = lowered
		}
		emojified := EmojiMap[emojiSequence]
		if emojified == "" {
			start = -1
			// Solves cases such as this ":sunglassesö:sunglasses:", where
			// the sequence wouldn't be sucessfully resolved otherwise.
			// Danke Marvin.
			index--
			continue
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

// toLower is an optimised variant of strings.ToLower. It only works for ASCII
// and returns an empty string if nothing has changed, this reduces return
// parameters, which in turn improves performance. It also avoids allocations
// where possible.
func toLower(input string) string {
	var out []byte
	for i := 0; i < len(input); i++ {
		c := input[i]
		if 'A' <= c && c <= 'Z' {
			if out == nil {
				out = []byte(input)
			}
			out[i] = c + 32
		}
	}

	if out != nil {
		return string(out)
	}
	return ""
}

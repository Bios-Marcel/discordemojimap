// Package discordemojimap provides a Replace function in order to escape
// emoji sequences with their respective emojis.
package discordemojimap

// Replace all emoji sequences contained in the emoji map with their
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
		// Single CodePoint UTF-8 is equal to ASCII 0-127. Meaning that
		// everything bigger than 128 is a multi-byte sequence. And no byte of
		// a multi-byte sequence is equal to 58 (ASCII for ":") or any other
		// byte for that matter.
		if input[index] != ':' {
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
			// If we have more than one emoji sequence, we might end up with a
			// buffer bigger than necessary, but that's fine.
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
		if c >= 'A' && c <= 'Z' {
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

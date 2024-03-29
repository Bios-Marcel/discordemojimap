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

		var buffer []byte
		if len(input) <= 32 {
			buffer = make([]byte, 0, 32)
		} else {
			// If we have more than one emoji sequence, we might end up with a
			// buffer bigger than necessary, but that's fine.
			buffer = make([]byte, 0, len(input)-len(emojiSequence)-2+len(emojified))
		}

		buffer = append(buffer, input[lastEnd:start]...)
		buffer = append(buffer, emojified...)
		lastEnd = index + 1

		// Prevent doing the checks we just did again
		index++

		// We use pretty much the same loop again to keep iterating from this
		// point tonward without reallocating all the time. This allows us to
		// stack allocate the buffer, as otherwise, the compiler won't do so.
		// Comments have been removed, as the logic is equivalent.
		for ; index < len(input); index++ {
			if input[index] != ':' {
				continue
			}

			if lastEnd > start || index-start == 1 {
				start = index
				continue
			}

			emojiSequence = input[start+1 : index]
			if lowered := toLower(emojiSequence); lowered != "" {
				emojiSequence = lowered
			}
			emojified = EmojiMap[emojiSequence]
			if emojified == "" {
				start = -1
				index--
				continue
			}

			buffer = append(buffer, input[lastEnd:start]...)
			buffer = append(buffer, emojified...)
			lastEnd = index + 1
		}

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
// and bound checks where possible.
//
//go:noinline While inlinable, inlining reduce performance. WHY? NO CLUE
func toLower(input string) string {
	for i := len(input) - 1; i >= 0; i-- {
		c := input[i]
		if c >= 'A' && c <= 'Z' {
			out := []byte(input)
			i = len(input) - 1
			// Eliminate further bound checks, since we iterate backwards, we
			// know that i >= 0 and our current i is the highest we'll check.
			_ = out[i]
			for ; i >= 0; i-- {
				c = input[i]
				if c >= 'A' && c <= 'Z' {
					out[i] = c + 32
				}
			}

			return string(out)
		}
	}
	return ""
}

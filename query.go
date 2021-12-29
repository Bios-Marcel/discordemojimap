package discordemojimap

import "strings"

// ContainsEmoji returns true if that emoji is mapped to one or more key.
func ContainsEmoji(emoji string) bool {
	for _, emojiInMap := range EmojiMap {
		if emojiInMap == emoji {
			return true
		}
	}

	return false
}

// ContainsCode returns true if emojiCode is mapped to an emoji.
func ContainsCode(emojiCode string) bool {
	_, contains := EmojiMap[emojiCode]
	return contains
}

// GetEmojiCodes contains all codes for an emoji in an array. If no code could
// be found, then the resulting array will be empty.
func GetEmojiCodes(emoji string) []string {
	var codes []string

	for code, emojiInMap := range EmojiMap {
		if emojiInMap == emoji {
			codes = append(codes, code)
		}
	}

	return codes
}

// GetEmoji returns the matching emoji or an empty string in case no match was
// found for the given code.
func GetEmoji(emojiCode string) string {
	return EmojiMap[emojiCode]
}

// GetEntriesWithPrefix returns a map of all found emojis with the given prefix.
//
// The function will search without accounting for trailing colons. As such, the
// caller should strip trailing colons, if any. The function, however, will
// search without case-sensitivity, as everything in the map is lower-cased.
//
// The function will return a nil map if prefix is empty.
func GetEntriesWithPrefix(prefix string) (matches map[string]string) {
	if prefix == "" {
		return nil
	}

	matches = make(map[string]string)
	prefix = strings.ToLower(prefix)

	for emojiCode, emoji := range EmojiMap {
		if strings.HasPrefix(emojiCode, prefix) {
			matches[emojiCode] = emoji
		}
	}

	return
}

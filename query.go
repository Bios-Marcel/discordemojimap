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

// ContainsCode returns true if emojiCode is mapped to an emoji. The search is
// case-insensitive.
func ContainsCode(emojiCode string) bool {
	if lowered := toLower(emojiCode); lowered != "" {
		emojiCode = lowered
	}
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
// The function will search without accounting for leading colons. The search
// is case-insensitive.
func GetEntriesWithPrefix(prefix string) map[string]string {
	matches := make(map[string]string)
	if prefix == "" {
		return matches
	}

	if lowered := toLower(prefix); lowered != "" {
		prefix = lowered
	}

	for emojiCode, emoji := range EmojiMap {
		if strings.HasPrefix(emojiCode, prefix) {
			matches[emojiCode] = emoji
		}
	}

	return matches
}

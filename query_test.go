package discordemojimap

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func TestContainsEmoji(t *testing.T) {
	if !ContainsEmoji("😀") {
		t.Error("Emoji '😀' should have been found.")
	}

	if ContainsEmoji("") {
		t.Error("An empty string should not have been found.")
	}

	if ContainsEmoji("OwO") {
		t.Error("An random string should not have been found.")
	}
}

func ExampleContainsEmoji() {
	fmt.Println(ContainsEmoji("😀"))
	// Output: true
}

func TestContainsEmojiCode(t *testing.T) {
	if !ContainsCode("grinning") {
		t.Error("The grinnign emoji '😀' should be in there as `grinning`.")
	}

	if ContainsEmoji("") {
		t.Error("An empty string should not have been found.")
	}

	if ContainsEmoji("incorrect code") {
		t.Error("An random string should not have been found.")
	}
}

func ExampleContainsCode() {
	fmt.Println(ContainsCode("grimacing"))
	// Output: true
}

func TestGetEmojiCodes(t *testing.T) {
	matches := GetEmojiCodes("🦁")
	if len(matches) != 2 {
		t.Error("There should have been two matches for '🦁'.")
	}

	if matches[0] != "lion" && matches[1] != "lion" {
		t.Errorf("None of the returned values was 'lion'. Result was: %v", matches)
	}

	if matches[0] != "lion_face" && matches[1] != "lion_face" {
		t.Errorf("None of the returned values was 'lion_face'. Result was: %v", matches)
	}

	matchesNoneExpected := GetEmojiCodes(" ")

	if len(matchesNoneExpected) != 0 {
		t.Errorf("Input should have been empty, but was: %v", matchesNoneExpected)
	}

	matchesNoneExpected = GetEmojiCodes("")

	if len(matchesNoneExpected) != 0 {
		t.Errorf("Input should have been empty, but was: %v", matchesNoneExpected)
	}

	matchesNoneExpected = GetEmojiCodes("Invalid input")

	if len(matchesNoneExpected) != 0 {
		t.Errorf("Input should have been empty, but was: %v", matchesNoneExpected)
	}
}

func ExampleGetEmojiCodes() {
	var codes = GetEmojiCodes("🦁")
	sort.Strings(codes)
	fmt.Println(codes)
	// Output: [lion lion_face]
}

func TestGetEntriesWithPrefix(t *testing.T) {
	lionTest(t, "lio")
	lionTest(t, "lion")

	matches := GetEntriesWithPrefix("")
	if len(matches) != 0 {
		t.Errorf("Matches should have been empty, but were: %v", matches)
	}

	matches = GetEntriesWithPrefix(" ")
	if len(matches) != 0 {
		t.Errorf("Matches should have been empty, but were: %v", matches)
	}
}

func ExampleGetEntriesWithPrefix() {
	fmt.Printf("%+v\n", GetEntriesWithPrefix("lio"))
	// Output: map[lion:🦁 lion_face:🦁]
}

func lionTest(t *testing.T, input string) {
	matches := GetEntriesWithPrefix(input)

	if len(matches) != 2 {
		t.Errorf("There should have been two matches for 'lio'.")
	}

	lionOne := matches["lion"]
	if lionOne != "🦁" {
		t.Errorf("The matches were expected to contain 'lion'.")
	}

	lionTwo := matches["lion_face"]
	if lionTwo != "🦁" {
		t.Errorf("The matches were expected to contain 'lion_face'.")
	}
}

var emojiCodeRegex = regexp.MustCompile("(?s):[a-zA-Z0-9_]+:")

// oldRegexReplace all emoji sequences contained in the discord emoji map with their
// respective emojis.
func oldRegexReplace(input string) string {
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

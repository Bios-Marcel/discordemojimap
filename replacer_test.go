package discordemojimap

import (
	"regexp"
	"strings"
	"testing"
)

func TestReplace(t *testing.T) {
	tests := []struct{ name, input, want string }{
		{"unnecessary colon before valid code", "::+1:", ":👍"},
		{"Just two double colons", "::", "::"},
		{"Just two double colons in the middle of a sentence", "What a :: world.", "What a :: world."},
		{"Escaping currently isn't possible", "I am sad \\:cry:", "I am sad \\😢"},
		{"No present emoji", "I am sad", "I am sad"},
		{"No valid emoji", "I am sad :cry", "I am sad :cry"},
		{"No valid emoji 2", "I am sad cry:", "I am sad cry:"},
		{"No valid emoji 3", "I am sad :crycry:", "I am sad :crycry:"},
		{"One valid emoji followed by an incomplete sequence", "I am sad :cry:cry:", "I am sad 😢cry:"},
		{"Simple in-sentence replacement", "I am sad :cry:", "I am sad 😢"},
		{"Simple in-sentence replacement, but single letter uppercase", "I am sad :cRy:", "I am sad 😢"},
		{"Simple in-sentence replacement, but all uppercase", "I am sad :CRY:", "I am sad 😢"},
		{"Simple single emoji replacement without text", ":cry:", "😢"},
		{"Two equal emojis next to eachother", ":cry::cry:", "😢😢"},
		{"Two equal emojis next to eachother with a spaces around", " :cry: :cry: ", " 😢 😢 "},
		{"Two different emojis next to eachother with a spaces around", " :cry: :angry: ", " 😢 😠 "},
		{"Unicode precision: rainbow flag emoji", "\U0001f3f3\uFE0F\u200D\U0001F308", "🏳️‍🌈"},
	}

	for _, tt := range tests {
		if got := Replace(tt.input); got != tt.want {
			t.Errorf("Replace() = %q, want %q", got, tt.want)
		}
	}
}

var sink string
var inputVariations = [][2]string{
	{"empty string", ""},
	{"just a colon", ":"},
	{"empty emoji sequence", "::"},
	{"invalid emoji sequence with invalid characters", ":sunglassesö:sunglasses:"},
	{"valid single letter emoji sequence", ":a:"},
	{"no emoji sequence", "Hello"},
	{"no emoji sequence, but single colon", "Hello :"},
	{"a long word", "abcdefghijklmnopqrstuvwxyz"},
	{"empty emoji sequence in middle of text", "What a :: world."},
	{"standalone invalid emoji sequence", ":invalidinvalid:"},
	{"invalid emoji sequence with space before and after", " :invalidinvalid: "},
	{"invalid emoji sequence with word before", "Hello :invalidinvalid:"},
	{"invalid emoji sequence with word after", ":invalidinvalid: Hello"},
	{"invalid emoji sequence with word before and after", "Hello :invalidinvalid: Hello"},
	{"very long string with invalid emoji sequence in the middle", strings.Repeat("a", 1000) + ":invalidinvalid:" + strings.Repeat("b", 1000)},
	{"very long string with valid emoji sequence in the middle", strings.Repeat("a", 1000) + ":sunglasses:" + strings.Repeat("b", 1000)},
	{"standalone valid emoji sequence", ":sunglasses:"},
	{"standalone valid uppercased emoji sequence", ":SUNGLASSES:"},
	{"valid emoji sequence with word before", "hello :sunglasses:"},
	{"valid emoji sequence with word before and single colon after", "Hello :sunglasses::"},
	{"valid emoji sequence with word before followed by single colon and more text", "Hello :sunglasses::lol"},
	{"valid emoji sequence with word after", ":sunglasses: hello"},
	{"two valid emoji sequences with space inbetween", ":sunglasses: :sunglasses:"},
	{"two valid emoji sequence with no space inbetween", ":sunglasses::sunglasses:"},
	{"two valid emoji sequence with word inbetween", ":sunglasses: hello :sunglasses:"},
	{"one mismatch", ":UPPER:"},
	{"one match upper", ":CRY:"},
	{"long with all kinds of cases", "I am :man_technologist: :extended_ascii_ý: :invalid_sequence: :umlautö: from :flag_for_turkey:. Tests are :thumbs_up:"},
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

func TestNewReplaceAndOldReplaceBehaveTheSame(t *testing.T) {
	for _, test := range inputVariations {
		a := oldRegexReplace(test[1])
		b := Replace(test[1])
		if a != b {
			t.Errorf("Regex - NonRegex: %s - %s", a, b)
		}
	}
}

func BenchmarkOldRegexReplace(b *testing.B) {
	var tmp string
	for _, test := range inputVariations {
		b.Run(test[0], func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				tmp = oldRegexReplace(test[1])
			}
		})
	}
	sink = tmp
}

func BenchmarkReplace(b *testing.B) {
	var tmp string
	for _, test := range inputVariations {
		b.Run(test[0], func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				tmp = Replace(test[1])
			}
		})
	}
	sink = tmp
}

func Test_toLower(t *testing.T) {
	inputs := [][2]string{
		{"A", "a"},
		{"a", ""},
		{"F", "f"},
		{"f", ""},
		{"Z", "z"},
		{"z", ""},
		{"1", ""},
		{":", ""},
		{" ", ""},
		{"", ""},
		{"ö", ""},
		// Only base ASCII is supported
		{"Ö", ""},
		{"👍", ""},
	}

	for _, input := range inputs {
		if got := toLower(input[0]); got != input[1] {
			t.Errorf("toLower(%q) = %q, want %q", input[0], got, input[1])
		}
	}
}

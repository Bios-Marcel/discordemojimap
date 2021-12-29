package discordemojimap

import (
	"testing"
)

func TestReplace(t *testing.T) {
	tests := []struct{ name, input, want string }{
		{"Just two double colons", "::", "::"},
		{"Just two double colons in the middle of a sentence", "What a :: world.", "What a :: world."},
		{"Escaping currently isn't possible", "I am sad \\:cry:", "I am sad \\😢"},
		{"No present emoji", "I am sad", "I am sad"},
		{"No valid emoji", "I am sad :cry", "I am sad :cry"},
		{"No valid emoji 2", "I am sad cry:", "I am sad cry:"},
		{"No valid emoji 3", "I am sad :crycry:", "I am sad :crycry:"},
		{"One valid emoji followed by an incomplete sequence", "I am sad :cry:cry:", "I am sad 😢cry:"},
		{"Simple insentence replacement", "I am sad :cry:", "I am sad 😢"},
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
var inputVariations = []string{
	"",
	":",
	"::",
	":a:",
	"Hello",
	"Hello :",
	"abcdefghijklmnopqrstuvwxyz",
	"What a :: world.",
	":invalidinvalid:",
	"Hello :invalidinvalid:",
	":invalidinvalid: Hello",
	"Hello :invalidinvalid: Hello",
	":sunglasses:",
	":SUNGLASSES:",
	"hello :sunglasses:",
	"Hello :sunglasses::",
	"Hello :sunglasses::lol",
	":sunglasses: hello",
	":sunglasses: :sunglasses:",
	":sunglasses::sunglasses:",
	":sunglasses: hello :sunglasses:",
}

func TestNewReplaceAndOldReplaceBehaveTheSame(t *testing.T) {
	for _, test := range inputVariations {
		a := oldRegexReplace(test)
		b := Replace(test)
		if a != b {
			t.Errorf("Regex - NonRegex: %s - %s", a, b)
		}
	}
}

func BenchmarkOldRegexReplace(b *testing.B) {
	var tmp string
	for _, test := range inputVariations {
		b.Run(test, func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				tmp = oldRegexReplace(test)
			}
		})
	}
	sink = tmp
}

func BenchmarkReplace(b *testing.B) {
	var tmp string
	for _, test := range inputVariations {
		b.Run(test, func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				tmp = Replace(test)
			}
		})
	}
	sink = tmp
}

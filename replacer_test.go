package discordemojimap

import (
	"fmt"
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

func ExampleReplace() {
	fmt.Println(Replace("Hello World :sun_with_face:"))
	// Output: Hello World 🌞
}

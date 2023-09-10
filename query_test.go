package discordemojimap

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsEmoji(t *testing.T) {
	tests := []struct {
		name  string
		emoji string
		want  bool
	}{
		{
			name:  "empty string",
			emoji: "",
			want:  false,
		},
		{
			name:  "incorrect emoji",
			emoji: "agfkbjasjnkfajnksf",
			want:  false,
		},
		{
			name:  "correct emoji",
			emoji: "😀",
			want:  true,
		},
		{
			name:  "correct emoji with skin tone",
			emoji: "👍🏻",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsEmoji(tt.emoji); got != tt.want {
				t.Errorf("ContainsEmoji() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleContainsEmoji() {
	fmt.Println(ContainsEmoji("😀"))
	// Output: true
}

func TestContainsCode(t *testing.T) {
	tests := []struct {
		name      string
		emojiCode string
		want      bool
	}{
		{
			name:      "empty string",
			emojiCode: "",
			want:      false,
		},
		{
			name:      "incorrect code",
			emojiCode: "agfkbjasjnkfajnksf",
			want:      false,
		},
		{
			name:      "correct code",
			emojiCode: "grimacing",
			want:      true,
		},
		{
			name:      "correct code with uppercase",
			emojiCode: "GRIMACING",
			want:      true,
		},
		{
			name:      "correct code with uppercase except first rune",
			emojiCode: "gRIMACING",
			want:      true,
		},
		{
			name:      "correct code with uppercase except last rune",
			emojiCode: "GRIMACINg",
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsCode(tt.emojiCode); got != tt.want {
				t.Errorf("ContainsCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleContainsCode() {
	fmt.Println(ContainsCode("grimacing"))
	// Output: true
}

func TestGetEmojiCodes(t *testing.T) {
	tests := []struct {
		emoji string
		want  []string
	}{
		{
			emoji: "",
		},
		{
			emoji: " ",
		},
		{
			emoji: "asdkfbakfabjnk",
		},
		{
			emoji: "🛝",
			want:  []string{"playground_slide"},
		},
		{
			emoji: "🦁",
			want:  []string{"lion", "lion_face"},
		},
		{
			emoji: "👍🏻",
			want:  []string{"+1_tone1", "thumbup_tone1", "thumbsup_tone1"},
		},
		{
			emoji: "👍",
			want:  []string{"+1", "thumbup", "thumbsup"},
		},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			if got := GetEmojiCodes(tt.emoji); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("GetEmojiCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleGetEmojiCodes() {
	codes := GetEmojiCodes("🦁")
	sort.Strings(codes)
	fmt.Println(codes)
	// Output: [lion lion_face]
}

func TestGetEntriesWithPrefix(t *testing.T) {
	tests := []struct {
		prefix []string
		want   map[string]string
	}{
		{
			prefix: []string{""},
			want:   map[string]string{},
		},
		{
			prefix: []string{"lio", "LIO", "lion", "LION", "lIO", "LIo"},
			want: map[string]string{
				"lion":      "🦁",
				"lion_face": "🦁",
			},
		},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			for _, prefix := range tt.prefix {
				if got := GetEntriesWithPrefix(prefix); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetEntriesWithPrefix() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func ExampleGetEntriesWithPrefix() {
	fmt.Printf("%+v\n", GetEntriesWithPrefix("lio"))
	// Output: map[lion:🦁 lion_face:🦁]
}

func TestGetEmoji(t *testing.T) {
	tests := []struct {
		emojiCode string
		want      string
	}{
		{
			emojiCode: "",
			want:      "",
		},
		{
			emojiCode: " ",
			want:      "",
		},
		{
			emojiCode: "lio",
			want:      "",
		},
		{
			emojiCode: "asdkfbakfabjnk",
			want:      "",
		},
		{
			emojiCode: "playground_slide",
			want:      "🛝",
		},
		{
			emojiCode: "lion",
			want:      "🦁",
		},
		{
			emojiCode: "LION",
			want:      "🦁",
		},
		{
			emojiCode: "lION",
			want:      "🦁",
		},
		{
			emojiCode: "LIOn",
			want:      "🦁",
		},
	}
	for index, tt := range tests {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			if got := GetEmoji(tt.emojiCode); got != tt.want {
				t.Errorf("GetEmoji() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleGetEmoji() {
	fmt.Println(GetEmoji("lion"))
	// Output: 🦁
}

package discordemojimap

import "testing"

func TestReplace(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "just two double colons",
			input: "::",
			want:  "::",
		},
		{
			name:  "just two double colons in the middle of a sentence",
			input: "What a :: world.",
			want:  "What a :: world.",
		},
		{
			name:  "escaping currently isn't possible",
			input: "I am sad \\:cry:",
			want:  "I am sad \\😢",
		},
		{
			name:  "no present emoji",
			input: "I am sad",
			want:  "I am sad",
		},
		{
			name:  "no valid emoji",
			input: "I am sad :cry",
			want:  "I am sad :cry",
		},
		{
			name:  "no valid emoji 2",
			input: "I am sad cry:",
			want:  "I am sad cry:",
		},
		{
			name:  "no valid emoji 3",
			input: "I am sad :crycry:",
			want:  "I am sad :crycry:",
		},
		{
			name:  "one valid emoji followed by an incomplete sequence",
			input: "I am sad :cry:cry:",
			want:  "I am sad 😢cry:",
		},
		{
			name:  "simple insentence replacement",
			input: "I am sad :cry:",
			want:  "I am sad 😢",
		},
		{
			name:  "simple single emoji replacement without text",
			input: ":cry:",
			want:  "😢",
		},
		{
			name:  "Two equal emojis next to eachother",
			input: ":cry::cry:",
			want:  "😢😢",
		},
		{
			name:  "Two equal emojis next to eachother with a spaces around",
			input: " :cry: :cry: ",
			want:  " 😢 😢 ",
		},
		{
			name:  "Two different emojis next to eachother with a spaces around",
			input: " :cry: :angry: ",
			want:  " 😢 😠 ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Replace(tt.input); got != tt.want {
				t.Errorf("Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

package vmask_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vmask"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestMaskerSecret(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		keepLen  int
		pattern  string
		expected string
	}{
		{
			name:     "default mask",
			input:    "1234567890",
			keepLen:  8,
			pattern:  "*****",
			expected: "12345678*****",
		},
		{
			name:     "custom keep length",
			input:    "1234567890",
			keepLen:  5,
			pattern:  "*****",
			expected: "12345*****",
		},
		{
			name:     "custom pattern",
			input:    "1234567890",
			keepLen:  8,
			pattern:  "###",
			expected: "12345678###",
		},
		{
			name:     "input shorter than keep length",
			input:    "123",
			keepLen:  8,
			pattern:  "*****",
			expected: "123*****",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			masker := (&vmask.MaskerSecret{}).
				WithKeepLen(tt.keepLen).
				WithMaskPattern(tt.pattern)

			result := masker.Marshal("", tt.input)

			vtest.Equal(t, tt.expected, result)
		})
	}
}

func TestMaskerSecretDefault(t *testing.T) {
	masker := (&vmask.MaskerSecret{}).Default()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "normal input",
			input:    "1234567890",
			expected: "12345678*****",
		},
		{
			name:     "short input",
			input:    "123",
			expected: "123*****",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := masker.Marshal("", tt.input)

			vtest.Equal(t, tt.expected, result)
		})
	}
}

func TestMaskSecret(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		keepLen  int
		pattern  string
		expected string
	}{
		{
			name:     "normal case",
			input:    "1234567890",
			keepLen:  8,
			pattern:  "*****",
			expected: "12345678*****",
		},
		{
			name:     "short input",
			input:    "123",
			keepLen:  8,
			pattern:  "*****",
			expected: "123*****",
		},
		{
			name:     "empty input",
			input:    "",
			keepLen:  8,
			pattern:  "*****",
			expected: "*****",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := vmask.MaskSecret(tt.input, tt.keepLen, tt.pattern)

			vtest.Equal(t, tt.expected, result)
		})
	}
}

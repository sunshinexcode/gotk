package vmask_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vmask"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestMaskerMapDefault(t *testing.T) {
	masker := (&vmask.MaskerMap{}).Default()

	tests := []struct {
		name     string
		input    map[string]any
		expected map[string]any
	}{
		{
			name: "basic masking",
			input: map[string]any{
				"password": "1234567890",
				"name":     "John Doe",
			},
			expected: map[string]any{
				"password": "12345678*****",
				"name":     "John Doe",
			},
		},
		{
			name: "nested map",
			input: map[string]any{
				"user": map[string]any{
					"password": "1234567890",
					"name":     "John Doe",
				},
			},
			expected: map[string]any{
				"user": map[string]any{
					"password": "12345678*****",
					"name":     "John Doe",
				},
			},
		},
		{
			name: "nested slice",
			input: map[string]any{
				"users": []any{
					map[string]any{
						"password": "1234567890",
						"name":     "John Doe",
					},
					map[string]any{
						"password": "0987654321",
						"name":     "Jane Doe",
					},
				},
			},
			expected: map[string]any{
				"users": []any{
					map[string]any{
						"password": "12345678*****",
						"name":     "John Doe",
					},
					map[string]any{
						"password": "09876543*****",
						"name":     "Jane Doe",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := masker.MaskMap(tt.input)

			vtest.Equal(t, tt.expected, result)
		})
	}
}

func TestMaskerMapCustom(t *testing.T) {
	tests := []struct {
		name     string
		masker   *vmask.MaskerMap
		input    map[string]any
		expected map[string]any
	}{
		{
			name: "custom keep length",
			masker: (&vmask.MaskerMap{}).Default().
				WithKeepLen(4),
			input: map[string]any{
				"password": "1234567890",
			},
			expected: map[string]any{
				"password": "1234*****",
			},
		},
		{
			name: "custom mask pattern",
			masker: (&vmask.MaskerMap{}).Default().
				WithMaskPattern("###"),
			input: map[string]any{
				"password": "1234567890",
			},
			expected: map[string]any{
				"password": "12345678###",
			},
		},
		{
			name: "custom sensitive keys",
			masker: (&vmask.MaskerMap{}).Default().
				WithSensitiveKeys([]string{"password", "token"}),
			input: map[string]any{
				"password": "1234567890",
				"token":    "abc123",
				"key":      "xyz789",
			},
			expected: map[string]any{
				"password": "12345678*****",
				"token":    "abc123*****",
				"key":      "xyz789",
			},
		},
		{
			name:   "short input",
			masker: (&vmask.MaskerMap{}).Default(),
			input: map[string]any{
				"password": "123",
			},
			expected: map[string]any{
				"password": "123*****",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.masker.MaskMap(tt.input)

			vtest.Equal(t, tt.expected, result)
		})
	}
}

func TestMaskMap(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]any
		expected map[string]any
	}{
		{
			name: "using default masker",
			input: map[string]any{
				"password": "1234567890",
				"name":     "John Doe",
			},
			expected: map[string]any{
				"password": "12345678*****",
				"name":     "John Doe",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := vmask.MaskMap(tt.input)

			vtest.Equal(t, tt.expected, result)
		})
	}
}

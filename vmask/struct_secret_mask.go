package vmask

import (
	"fmt"

	masker "github.com/ggwhite/go-masker/v2"
)

var (
	// Mask is the global masker instance for handling secret masking
	Mask = masker.NewMaskerMarshaler()
)

func init() {
	// Register with default configuration
	RegisterSecretMaskers(5, 24, "*****")
}

// MaskerSecret implements secret masking with configurable prefix preservation
//
// The masker replaces sensitive content with a pattern while keeping a prefix:
// - By default keeps 5 characters (when using "secret" mask type)
// - Supports variable length prefixes using "secret:N" format (5 ≤ N ≤ 12)
//
// Examples:
//
//	MaskType: "secret"    -> Keep 5 chars + mask pattern (e.g. "12345*****")
//	MaskType: "secret:8"  -> Keep 8 chars + mask pattern (e.g. "12345678*****")
//
// Note: If input length is shorter than KeepLen, shows full available prefix
type MaskerSecret struct {
	KeepLen     int    // Number of leading characters to preserve
	MaskPattern string // Mask pattern to append after preserved characters
}

// Default returns a MaskerSecret with default values
// Default configuration:
// - KeepLen: 8 characters
// - MaskPattern: "*****"
func (m *MaskerSecret) Default() *MaskerSecret {
	m.KeepLen = MaskKeepLen
	m.MaskPattern = MaskPattern
	return m
}

// Marshal implements the masker.Marshaler interface
// It masks the input string while preserving the specified number of leading characters
func (m *MaskerSecret) Marshal(s string, val string) string {
	return MaskSecret(val, m.KeepLen, m.MaskPattern)
}

// WithKeepLen sets the number of characters to keep from the beginning of the string
// Returns the MaskerSecret instance for method chaining
func (m *MaskerSecret) WithKeepLen(keepLen int) *MaskerSecret {
	m.KeepLen = keepLen
	return m
}

// WithMaskPattern sets the pattern to use for masking
// Returns the MaskerSecret instance for method chaining
func (m *MaskerSecret) WithMaskPattern(pattern string) *MaskerSecret {
	m.MaskPattern = pattern
	return m
}

// MaskSecret masks a string while preserving a specified number of leading characters
//
// val:         the input string to mask
// keepLen:     number of leading characters to preserve
// maskPattern: the pattern to use for masking
//
// Return:
//   - If input length is less than keepLen, preserves all available characters
//   - Otherwise, preserves keepLen characters and appends the mask pattern
func MaskSecret(val string, keepLen int, maskPattern string) string {
	lens := len(val)
	if lens < keepLen {
		keepLen = lens
	}

	return val[:keepLen] + maskPattern
}

// RegisterSecretMaskers registers secret maskers with configurable options
//
// minLen:  minimum number of characters to keep (inclusive)
// maxLen:  maximum number of characters to keep (inclusive)
// pattern: the pattern to use for masking
//
// This function registers:
// 1. A default masker with 8 characters preserved
// 2. Additional maskers for each length from minLen to maxLen
func RegisterSecretMaskers(minLen, maxLen int, pattern string) {
	// Register default secret masker
	Mask.Register("secret", (&MaskerSecret{}).Default())

	// Register secret maskers with different keep lengths
	for i := minLen; i <= maxLen; i++ {
		Mask.Register(masker.MaskerType(fmt.Sprintf("secret:%d", i)), (&MaskerSecret{}).
			WithKeepLen(i).
			WithMaskPattern(pattern))
	}
}

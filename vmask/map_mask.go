package vmask

import "strings"

// Mask is the global instance of MaskerMap with default configuration
var MaskerMapDefault = (&MaskerMap{}).Default()

// MaskMap masks sensitive information in a map using the default configuration
func MaskMap(input map[string]any) map[string]any {
	return MaskerMapDefault.MaskMap(input)
}

// MaskerMap represents a configurable map masking utility
// It provides methods to mask sensitive information in maps and slices
type MaskerMap struct {
	// KeepLen specifies how many characters to keep from the beginning of sensitive values
	KeepLen int
	// MaskPattern is the pattern used to replace sensitive information
	MaskPattern string
	// SensitiveKeys contains keywords that indicate sensitive fields
	SensitiveKeys []string
}

// containsSensitiveKey checks if a key contains any sensitive keywords
func (m *MaskerMap) containsSensitiveKey(key string) bool {
	lowerKey := strings.ToLower(key)
	for _, sk := range m.SensitiveKeys {
		if strings.Contains(lowerKey, sk) {
			return true
		}
	}

	return false
}

// Default returns a MaskerMap with default configuration
func (m *MaskerMap) Default() *MaskerMap {
	m.SensitiveKeys = []string{"api_key", "key", "token", "password", "secret"}
	m.KeepLen = MaskKeepLen
	m.MaskPattern = MaskPattern
	return m
}

// MaskMap masks sensitive information in a map using the current configuration
// It processes nested maps and slices recursively
func (m *MaskerMap) MaskMap(input map[string]any) map[string]any {
	return m.maskSensitiveMap(input)
}

// maskSensitiveMap is the internal implementation for masking maps
func (m *MaskerMap) maskSensitiveMap(input map[string]any) map[string]any {
	masked := make(map[string]any)
	for k, v := range input {
		if m.containsSensitiveKey(k) {
			switch val := v.(type) {
			case string:
				lens := len(val)
				keepLen := m.KeepLen
				if lens < keepLen {
					keepLen = lens
				}
				masked[k] = val[:keepLen] + m.MaskPattern
			default:
				masked[k] = v
			}
			continue
		}

		switch val := v.(type) {
		case map[string]any:
			masked[k] = m.maskSensitiveMap(val)
		case []any:
			masked[k] = m.maskSensitiveSlice(val)
		default:
			masked[k] = v
		}
	}

	return masked
}

// maskSensitiveSlice is the internal implementation for masking slices
func (m *MaskerMap) maskSensitiveSlice(input []any) []any {
	masked := make([]any, len(input))
	for i, v := range input {
		if mapVal, ok := v.(map[string]any); ok {
			masked[i] = m.maskSensitiveMap(mapVal)
		} else {
			masked[i] = v
		}
	}
	return masked
}

// WithKeepLen sets the number of characters to keep from sensitive values
// Returns the MaskerMap instance for method chaining
func (m *MaskerMap) WithKeepLen(keepLen int) *MaskerMap {
	m.KeepLen = keepLen
	return m
}

// WithMaskPattern sets the pattern to use for masking
// Returns the MaskerMap instance for method chaining
func (m *MaskerMap) WithMaskPattern(pattern string) *MaskerMap {
	m.MaskPattern = pattern
	return m
}

// WithSensitiveKeys sets the sensitive keywords to detect
// Returns the MaskerMap instance for method chaining
func (m *MaskerMap) WithSensitiveKeys(keys []string) *MaskerMap {
	m.SensitiveKeys = keys
	return m
}

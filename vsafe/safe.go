package vsafe

import (
	"regexp"
)

// MaskPassword hide password
func MaskPassword(data string) string {
	return ReplaceData(data, `("Password":"|"Secret":")(.*)(".*)`, "$1***$3")
}

// MaskUrl hide sensitivity information for url
func MaskUrl(data string) string {
	return ReplaceData(data, `(http|https|mongodb://)(.*:)(.*)(@.*)`, "$1$2***$4")
}

// ReplaceData replaces data
func ReplaceData(data string, reg string, repl string) string {
	return regexp.MustCompile(reg).ReplaceAllString(data, repl)
}

package vstr

import (
	"fmt"

	"github.com/gogf/gf/v2/text/gstr"
)

// Equal reports whether `a` and `b`, interpreted as UTF-8 strings,
// are equal under Unicode case-folding, case-insensitively.
func Equal(a, b string) bool {
	return gstr.Equal(a, b)
}

// Repeat returns a new string consisting of multiplier copies of the string input.
func Repeat(input string, multiplier int) string {
	return gstr.Repeat(input, multiplier)
}

// S shortcut for fmt.Sprintf
func S(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// StrLimit returns a portion of string `str` specified by `length` parameters, if the length
// of `str` is greater than `length`, then the `suffix` will be appended to the result string.
// considers parameter `str` as unicode string.
func StrLimit(str string, length int, suffix ...string) string {
	return gstr.StrLimitRune(str, length, suffix...)
}

// SubStr returns a portion of string `str` specified by the `start` and `length` parameters.
// considers parameter `str` as unicode string.
// The parameter `length` is optional, it uses the length of `str` in default.
func SubStr(str string, start int, length ...int) (substr string) {
	return gstr.SubStrRune(str, start, length...)
}

// Trim strips whitespace (or other characters) from the beginning and end of a string.
// The optional parameter `characterMask` specifies the additional stripped characters.
func Trim(str string, characterMask ...string) string {
	return gstr.Trim(str, characterMask...)
}

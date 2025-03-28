package vbase64

import (
	"github.com/gogf/gf/v2/encoding/gbase64"
)

// Decode decodes bytes with BASE64 algorithm.
func Decode(data []byte) ([]byte, error) {
	return gbase64.Decode(data)
}

// DecodeStr decodes string with BASE64 algorithm.
func DecodeStr(data string) ([]byte, error) {
	return gbase64.DecodeString(data)
}

// DecodeToStr decodes string with BASE64 algorithm.
func DecodeToStr(data string) (string, error) {
	return gbase64.DecodeToString(data)
}

// Encode encodes bytes with BASE64 algorithm.
func Encode(src []byte) []byte {
	return gbase64.Encode(src)
}

// EncodeStr encodes string with BASE64 algorithm.
func EncodeStr(src string) string {
	return gbase64.EncodeString(src)
}

// EncodeToStr encodes bytes to string with BASE64 algorithm.
func EncodeToStr(src []byte) string {
	return gbase64.EncodeToString(src)
}

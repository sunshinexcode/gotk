package vmd5

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

// Get encrypts any type of variable using MD5 algorithms.
// It uses gconv package to convert `v` to its bytes type.
func Get(data any) (encrypt string, err error) {
	return gmd5.Encrypt(data)
}

// GetByte get byte
func GetByte(data string) []byte {
	h := md5.New()
	h.Write([]byte(data))

	return h.Sum(nil)
}

// GetStr get string, lowercase
func GetStr(data string) string {
	h := md5.New()
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}

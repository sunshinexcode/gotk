package vaes

import (
	"github.com/forgoer/openssl"
)

// DecryptCBC decrypts `cipherText` using CBC mode.
// Note that the key must be 16/24/32 bit length.
func DecryptCBC(src, key, iv []byte, padding string) ([]byte, error) {
	return openssl.AesCBCDecrypt(src, key, iv, padding)
}

// EncryptCBC encrypts `src` using CBC mode.
// Note that the key must be 16/24/32 bit length.
func EncryptCBC(src, key, iv []byte, padding string) ([]byte, error) {
	return openssl.AesCBCEncrypt(src, key, iv, padding)
}

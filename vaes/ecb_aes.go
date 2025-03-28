package vaes

import (
	"github.com/forgoer/openssl"
)

// DecryptECB decrypts `cipherText` using ECB mode.
func DecryptECB(src, key []byte, padding string) ([]byte, error) {
	return openssl.AesECBDecrypt(src, key, padding)
}

// EncryptECB encrypts `src` using ECB mode.
func EncryptECB(src, key []byte, padding string) ([]byte, error) {
	return openssl.AesECBEncrypt(src, key, padding)
}

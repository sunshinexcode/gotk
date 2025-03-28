package vaes

import (
	"crypto/aes"
	"crypto/cipher"
)

// DecryptGCM decrypts `cipherText` using GCM mode.
func DecryptGCM(cipherText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, _ := cipher.NewGCM(block)
	plainText, err := gcm.Open(nil, iv, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

// EncryptGCM encrypts `plainText` using GCM mode.
func EncryptGCM(plainText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, _ := cipher.NewGCM(block)
	cipherText := gcm.Seal(nil, iv, plainText, nil)
	return cipherText, nil
}

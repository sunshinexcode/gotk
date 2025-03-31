package vaes

import (
	"crypto/aes"
	"crypto/cipher"
)

// DecryptGcm decrypts data using AES in GCM mode
//
// cipherText: the encrypted data to decrypt
// key:        the encryption key (must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256)
// iv:         the nonce/initialization vector (must be 12 bytes for GCM mode)
//
// Return: decrypted data as byte slice. error if decryption fails
//
// Example:
//
//	key := []byte("1234567890123456") // 16 bytes for AES-128
//	iv := []byte("123456789012")      // 12 bytes for GCM
//	ciphertext := []byte("...")
//	plaintext, err := DecryptGcm(ciphertext, key, iv)
func DecryptGcm(cipherText []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plainText, err := gcm.Open(nil, iv, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

// EncryptGcm encrypts data using AES in GCM mode
//
// plainText: the data to encrypt
// key:       the encryption key (must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256)
// iv:        the nonce/initialization vector (must be 12 bytes for GCM mode)
//
// Return: encrypted data as byte slice. error if encryption fails
//
// Example:
//
//	key := []byte("1234567890123456") // 16 bytes for AES-128
//	iv := []byte("123456789012")      // 12 bytes for GCM
//	plaintext := []byte("Hello, World!")
//	ciphertext, err := EncryptGcm(plaintext, key, iv)
func EncryptGcm(plainText []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nil, iv, plainText, nil)
	return cipherText, nil
}

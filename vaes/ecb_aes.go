package vaes

import (
	"github.com/forgoer/openssl"
)

// DecryptEcb decrypts data using AES in ECB mode
//
// src:     the ciphertext to decrypt
// key:     the encryption key (must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256)
// padding: the padding method
//
// Return: decrypted data as byte slice. error if decryption fails
//
// Example:
//
//	key := []byte("1234567890123456") // 16 bytes for AES-128
//	ciphertext := []byte("...")
//	plaintext, err := DecryptEcb(ciphertext, key, "PKCS7")
func DecryptEcb(src []byte, key []byte, padding string) ([]byte, error) {
	return openssl.AesECBDecrypt(src, key, padding)
}

// EncryptEcb encrypts data using AES in ECB mode
//
// src:     the plaintext to encrypt
// key:     the encryption key (must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256)
// padding: the padding method
//
// Return: encrypted data as byte slice. error if encryption fails
//
// Example:
//
//	key := []byte("1234567890123456") // 16 bytes for AES-128
//	plaintext := []byte("Hello, World!")
//	ciphertext, err := EncryptEcb(plaintext, key, "PKCS7")
func EncryptEcb(src []byte, key []byte, padding string) ([]byte, error) {
	return openssl.AesECBEncrypt(src, key, padding)
}

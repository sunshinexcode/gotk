package vaes

import (
	"github.com/forgoer/openssl"
)

// DecryptCbc decrypts data using AES in CBC mode
//
// src:     the ciphertext to decrypt
// key:     the encryption key (must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256)
// iv:      the initialization vector (must be 16 bytes)
// padding: the padding method
//
// Return: decrypted data as byte slice. error if decryption fails
//
// Example:
//
//	key := []byte("1234567890123456") // 16 bytes for AES-128
//	iv := []byte("1234567890123456")  // 16 bytes
//	ciphertext := []byte("...")
//	plaintext, err := DecryptCbc(ciphertext, key, iv, "PKCS7")
func DecryptCbc(src []byte, key []byte, iv []byte, padding string) ([]byte, error) {
	return openssl.AesCBCDecrypt(src, key, iv, padding)
}

// EncryptCbc encrypts data using AES in CBC mode
//
// src:     the plaintext to encrypt
// key:     the encryption key (must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256)
// iv:      the initialization vector (must be 16 bytes)
// padding: the padding method
//
// Return: encrypted data as byte slice. error if encryption fails
//
// Example:
//
//	key := []byte("1234567890123456") // 16 bytes for AES-128
//	iv := []byte("1234567890123456")  // 16 bytes
//	plaintext := []byte("Hello, World!")
//	ciphertext, err := EncryptCbc(plaintext, key, iv, "PKCS7")
func EncryptCbc(src []byte, key []byte, iv []byte, padding string) ([]byte, error) {
	return openssl.AesCBCEncrypt(src, key, iv, padding)
}

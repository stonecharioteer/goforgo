// encryption_aes.go
// Learn AES encryption and decryption in Go

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	fmt.Println("=== AES Encryption/Decryption ===")
	
	// TODO: Generate a random key
	key := make([]byte, 32) // AES-256 key
	// Fill key with random bytes
	
	plaintext := "This is a secret message that needs to be encrypted!"
	fmt.Printf("Original text: %s\n", plaintext)
	
	// TODO: Create AES cipher
	block, err := /* create AES cipher with key */
	if err != nil {
		fmt.Printf("Error creating cipher: %v\n", err)
		return
	}
	
	fmt.Printf("AES block size: %d bytes\n", block.BlockSize())
	
	fmt.Println("\n=== AES-GCM Encryption ===")
	
	// TODO: Create GCM mode
	gcm, err := /* create GCM mode from block cipher */
	if err != nil {
		fmt.Printf("Error creating GCM: %v\n", err)
		return
	}
	
	// TODO: Generate random nonce
	nonce := make([]byte, gcm.NonceSize())
	// Fill nonce with random bytes
	
	fmt.Printf("Nonce size: %d bytes\n", len(nonce))
	fmt.Printf("Nonce: %x\n", nonce)
	
	// TODO: Encrypt the plaintext
	ciphertext := /* encrypt plaintext using GCM */
	
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	fmt.Printf("Ciphertext length: %d bytes\n", len(ciphertext))
	
	fmt.Println("\n=== AES-GCM Decryption ===")
	
	// TODO: Decrypt the ciphertext
	decrypted, err := /* decrypt ciphertext using GCM */
	if err != nil {
		fmt.Printf("Decryption error: %v\n", err)
		return
	}
	
	fmt.Printf("Decrypted text: %s\n", string(decrypted))
	
	// TODO: Verify decryption was successful
	if /* compare original and decrypted */ {
		fmt.Println("✓ Decryption successful!")
	} else {
		fmt.Println("✗ Decryption failed!")
	}
	
	fmt.Println("\n=== AES-CBC Encryption ===")
	
	// TODO: Implement PKCS7 padding function
	pad := func(data []byte, blockSize int) []byte {
		// Calculate padding needed
		// Add padding bytes
		// Return padded data
	}
	
	// TODO: Implement PKCS7 unpadding function
	unpad := func(data []byte) []byte {
		// Get padding length from last byte
		// Remove padding
		// Return unpadded data
	}
	
	// TODO: Create CBC mode
	plaintextBytes := []byte(plaintext)
	
	// Pad the plaintext
	paddedText := /* pad plaintext to block size */
	
	// Generate IV
	iv := make([]byte, aes.BlockSize)
	// Fill IV with random bytes
	
	fmt.Printf("IV: %x\n", iv)
	
	// TODO: Create CBC encrypter
	mode := /* create CBC encrypter */
	
	// TODO: Encrypt using CBC
	cbcCiphertext := make([]byte, len(paddedText))
	// Encrypt paddedText
	
	fmt.Printf("CBC Ciphertext: %x\n", cbcCiphertext)
	
	fmt.Println("\n=== AES-CBC Decryption ===")
	
	// TODO: Create CBC decrypter
	decMode := /* create CBC decrypter */
	
	// TODO: Decrypt using CBC
	cbcDecrypted := make([]byte, len(cbcCiphertext))
	// Decrypt ciphertext
	
	// Remove padding
	finalDecrypted := /* unpad the decrypted data */
	
	fmt.Printf("CBC Decrypted: %s\n", string(finalDecrypted))
	
	fmt.Println("\n=== Key Derivation Demo ===")
	
	// TODO: Derive key from password (simple example)
	password := "my-secret-password"
	// In production, use PBKDF2, scrypt, or Argon2
	
	// Simple key derivation (NOT secure for production)
	derivedKey := /* derive key from password using SHA256 */
	
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Derived key: %x\n", derivedKey)
	
	// TODO: Use derived key for encryption
	derivedBlock, err := /* create AES cipher with derived key */
	if err != nil {
		fmt.Printf("Error with derived key: %v\n", err)
		return
	}
	
	fmt.Println("✓ Successfully created cipher with derived key")
	
	fmt.Println("\n=== Security Notes ===")
	fmt.Println("• Always use cryptographically secure random numbers")
	fmt.Println("• Never reuse nonces/IVs with the same key")
	fmt.Println("• Use proper key derivation functions for passwords")
	fmt.Println("• Consider using higher-level libraries like NaCl/Box")
	fmt.Println("• AES-GCM provides both encryption and authentication")
}
// encryption_aes.go - SOLUTION
// Learn AES encryption and decryption in Go

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	fmt.Println("=== AES Encryption/Decryption ===")
	
	// Generate a random key
	key := make([]byte, 32) // AES-256 key
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		fmt.Printf("Error generating key: %v\n", err)
		return
	}
	
	plaintext := "This is a secret message that needs to be encrypted!"
	fmt.Printf("Original text: %s\n", plaintext)
	
	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("Error creating cipher: %v\n", err)
		return
	}
	
	fmt.Printf("AES block size: %d bytes\n", block.BlockSize())
	
	fmt.Println("\n=== AES-GCM Encryption ===")
	
	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Printf("Error creating GCM: %v\n", err)
		return
	}
	
	// Generate random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Printf("Error generating nonce: %v\n", err)
		return
	}
	
	fmt.Printf("Nonce size: %d bytes\n", len(nonce))
	fmt.Printf("Nonce: %x\n", nonce)
	
	// Encrypt the plaintext
	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)
	
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	fmt.Printf("Ciphertext length: %d bytes\n", len(ciphertext))
	
	fmt.Println("\n=== AES-GCM Decryption ===")
	
	// Decrypt the ciphertext
	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Printf("Decryption error: %v\n", err)
		return
	}
	
	fmt.Printf("Decrypted text: %s\n", string(decrypted))
	
	// Verify decryption was successful
	if string(decrypted) == plaintext {
		fmt.Println("✓ Decryption successful!")
	} else {
		fmt.Println("✗ Decryption failed!")
	}
	
	fmt.Println("\n=== Key Derivation Demo ===")
	
	// Derive key from password (simple example)
	password := "my-secret-password"
	// In production, use PBKDF2, scrypt, or Argon2
	
	// Simple key derivation (NOT secure for production)
	derivedKey := sha256.Sum256([]byte(password))
	
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Derived key: %x\n", derivedKey)
	
	// Use derived key for encryption
	derivedBlock, err := aes.NewCipher(derivedKey[:])
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
// hashing_basics.go
// Learn cryptographic hashing with Go's crypto package

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func main() {
	fmt.Println("=== Cryptographic Hashing ===")
	
	data := "Hello, Go Cryptography!"
	
	// TODO: Create MD5 hash
	md5Hash := md5.New()
	// Write data to hash
	// Get the final hash sum
	fmt.Printf("MD5: %x\n", /* get hash result */)
	
	// TODO: Create SHA1 hash
	sha1Hash := /* create SHA1 hasher */
	// Write data and get result
	fmt.Printf("SHA1: %x\n", /* get hash result */)
	
	// TODO: Create SHA256 hash
	sha256Hash := /* create SHA256 hasher */
	// Write data and get result
	fmt.Printf("SHA256: %x\n", /* get hash result */)
	
	// TODO: Create SHA512 hash
	sha512Hash := /* create SHA512 hasher */
	// Write data and get result
	fmt.Printf("SHA512: %x\n", /* get hash result */)
	
	fmt.Println("\n=== Convenient Hash Functions ===")
	
	// TODO: Use convenience functions
	dataBytes := []byte(data)
	
	md5Sum := /* use md5.Sum() function */
	sha1Sum := /* use sha1.Sum() function */
	sha256Sum := /* use sha256.Sum256() function */
	sha512Sum := /* use sha512.Sum512() function */
	
	fmt.Printf("MD5 (direct): %x\n", md5Sum)
	fmt.Printf("SHA1 (direct): %x\n", sha1Sum)
	fmt.Printf("SHA256 (direct): %x\n", sha256Sum)
	fmt.Printf("SHA512 (direct): %x\n", sha512Sum)
	
	fmt.Println("\n=== Password Hashing Comparison ===")
	
	passwords := []string{"password123", "admin", "secretkey", "password123"}
	
	fmt.Println("Password\t\tSHA256 Hash")
	fmt.Println("--------\t\t-----------")
	
	for _, pwd := range passwords {
		// TODO: Hash each password with SHA256
		hash := /* hash the password */
		fmt.Printf("%-15s\t%x\n", pwd, hash)
	}
	
	fmt.Println("\n=== File-like Hashing ===")
	
	// TODO: Hash large data incrementally
	hasher := sha256.New()
	
	// Simulate reading data in chunks
	chunks := []string{
		"This is chunk 1. ",
		"This is chunk 2. ",
		"This is the final chunk.",
	}
	
	for i, chunk := range chunks {
		// Write each chunk to hasher
		// Print intermediate state info
		fmt.Printf("Added chunk %d (size: %d bytes)\n", i+1, len(chunk))
	}
	
	finalHash := /* get final hash */
	fmt.Printf("Final hash: %x\n", finalHash)
	
	fmt.Println("\n=== Hash Verification ===")
	
	// TODO: Verify data integrity
	originalData := "Important data that must not be tampered with"
	originalHash := /* create hash of original data */
	
	// Simulate data that might be corrupted
	testData := []string{
		"Important data that must not be tampered with",  // Same
		"Important data that must not be tampered with.", // Extra period
		"important data that must not be tampered with",  // Different case
		"Important data that must not be tampered with!", // Different punctuation
	}
	
	fmt.Printf("Original hash: %x\n", originalHash)
	fmt.Println("\nVerification results:")
	
	for i, test := range testData {
		// TODO: Hash test data and compare
		testHash := /* hash test data */
		isValid := /* compare hashes */
		
		status := "✓ VALID"
		if !isValid {
			status = "✗ INVALID"
		}
		
		fmt.Printf("Test %d: %s - %s\n", i+1, status, test)
	}
}
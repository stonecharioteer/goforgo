// hashing_basics.go - SOLUTION
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
	
	// Create MD5 hash
	md5Hash := md5.New()
	io.WriteString(md5Hash, data)
	fmt.Printf("MD5: %x\n", md5Hash.Sum(nil))
	
	// Create SHA1 hash
	sha1Hash := sha1.New()
	io.WriteString(sha1Hash, data)
	fmt.Printf("SHA1: %x\n", sha1Hash.Sum(nil))
	
	// Create SHA256 hash
	sha256Hash := sha256.New()
	io.WriteString(sha256Hash, data)
	fmt.Printf("SHA256: %x\n", sha256Hash.Sum(nil))
	
	// Create SHA512 hash
	sha512Hash := sha512.New()
	io.WriteString(sha512Hash, data)
	fmt.Printf("SHA512: %x\n", sha512Hash.Sum(nil))
	
	fmt.Println("\n=== Convenient Hash Functions ===")
	
	// Use convenience functions
	dataBytes := []byte(data)
	
	md5Sum := md5.Sum(dataBytes)
	sha1Sum := sha1.Sum(dataBytes)
	sha256Sum := sha256.Sum256(dataBytes)
	sha512Sum := sha512.Sum512(dataBytes)
	
	fmt.Printf("MD5 (direct): %x\n", md5Sum)
	fmt.Printf("SHA1 (direct): %x\n", sha1Sum)
	fmt.Printf("SHA256 (direct): %x\n", sha256Sum)
	fmt.Printf("SHA512 (direct): %x\n", sha512Sum)
	
	fmt.Println("\n=== Password Hashing Comparison ===")
	
	passwords := []string{"password123", "admin", "secretkey", "password123"}
	
	fmt.Println("Password\t\tSHA256 Hash")
	fmt.Println("--------\t\t-----------")
	
	for _, pwd := range passwords {
		hash := sha256.Sum256([]byte(pwd))
		fmt.Printf("%-15s\t%x\n", pwd, hash)
	}
	
	fmt.Println("\n=== File-like Hashing ===")
	
	// Hash large data incrementally
	hasher := sha256.New()
	
	// Simulate reading data in chunks
	chunks := []string{
		"This is chunk 1. ",
		"This is chunk 2. ",
		"This is the final chunk.",
	}
	
	for i, chunk := range chunks {
		hasher.Write([]byte(chunk))
		fmt.Printf("Added chunk %d (size: %d bytes)\n", i+1, len(chunk))
	}
	
	finalHash := hasher.Sum(nil)
	fmt.Printf("Final hash: %x\n", finalHash)
	
	fmt.Println("\n=== Hash Verification ===")
	
	// Verify data integrity
	originalData := "Important data that must not be tampered with"
	originalHash := sha256.Sum256([]byte(originalData))
	
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
		testHash := sha256.Sum256([]byte(test))
		isValid := testHash == originalHash
		
		status := "✓ VALID"
		if !isValid {
			status = "✗ INVALID"
		}
		
		fmt.Printf("Test %d: %s - %s\n", i+1, status, test)
	}
}
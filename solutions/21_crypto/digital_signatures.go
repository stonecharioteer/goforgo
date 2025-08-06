// digital_signatures.go - SOLUTION
// Learn RSA and ECDSA digital signatures for authentication and non-repudiation

package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Digital Signatures ===")
	
	fmt.Println("\n=== RSA Digital Signatures ===")
	
	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Error generating RSA key: %v\n", err)
		return
	}
	
	publicKey := &privateKey.PublicKey
	
	fmt.Printf("Generated RSA key pair (size: %d bits)\n", privateKey.Size()*8)
	
	// Message to sign
	message := []byte("This is a message to be digitally signed")
	fmt.Printf("Original message: %s\n", message)
	
	// Create message hash
	hash := sha256.Sum256(message)
	
	// Sign the hash using RSA private key
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hash[:], nil)
	if err != nil {
		fmt.Printf("Error signing message: %v\n", err)
		return
	}
	
	fmt.Printf("Signature created (length: %d bytes)\n", len(signature))
	
	// Verify the signature using RSA public key
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, hash[:], signature, nil)
	if err != nil {
		fmt.Printf("❌ Signature verification failed: %v\n", err)
	} else {
		fmt.Println("✅ RSA signature verification successful")
	}
	
	// Test with tampered message
	tamperedMessage := []byte("This is a TAMPERED message to be digitally signed")
	tamperedHash := sha256.Sum256(tamperedMessage)
	
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, tamperedHash[:], signature, nil)
	if err != nil {
		fmt.Println("✅ Tampered message correctly rejected")
	} else {
		fmt.Println("❌ Tampered message incorrectly accepted")
	}
	
	fmt.Println("\n=== ECDSA Digital Signatures ===")
	
	// Generate ECDSA key pair using P-256 curve
	ecPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Printf("Error generating ECDSA key: %v\n", err)
		return
	}
	
	ecPublicKey := &ecPrivateKey.PublicKey
	
	fmt.Printf("Generated ECDSA key pair (curve: %s)\n", ecPrivateKey.Curve.Params().Name)
	
	// Sign message with ECDSA
	ecMessage := []byte("ECDSA message for signing")
	ecHash := sha256.Sum256(ecMessage)
	
	// Sign the hash using ECDSA
	r, s, err := ecdsa.Sign(rand.Reader, ecPrivateKey, ecHash[:])
	if err != nil {
		fmt.Printf("Error signing with ECDSA: %v\n", err)
		return
	}
	
	fmt.Printf("ECDSA signature created (r: %d bytes, s: %d bytes)\n", 
		len(r.Bytes()), len(s.Bytes()))
	
	// Verify ECDSA signature
	valid := ecdsa.Verify(ecPublicKey, ecHash[:], r, s)
	if valid {
		fmt.Println("✅ ECDSA signature verification successful")
	} else {
		fmt.Println("❌ ECDSA signature verification failed")
	}
	
	// Test ECDSA with tampered message
	ecTamperedMessage := []byte("ECDSA TAMPERED message for signing")
	ecTamperedHash := sha256.Sum256(ecTamperedMessage)
	
	valid = ecdsa.Verify(ecPublicKey, ecTamperedHash[:], r, s)
	if !valid {
		fmt.Println("✅ ECDSA tampered message correctly rejected")
	} else {
		fmt.Println("❌ ECDSA tampered message incorrectly accepted")
	}
	
	fmt.Println("\n=== Digital Signature Comparison ===")
	
	// Compare RSA vs ECDSA performance and signature sizes
	testMessage := []byte("Performance test message for signature comparison")
	testHash := sha256.Sum256(testMessage)
	
	// Time RSA signing
	startTime := time.Now()
	rsaSignature, _ := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, testHash[:], nil)
	rsaSignTime := time.Since(startTime)
	
	// Time ECDSA signing  
	startTime = time.Now()
	ecR, ecS, _ := ecdsa.Sign(rand.Reader, ecPrivateKey, testHash[:])
	ecSignTime := time.Since(startTime)
	
	// Calculate signature sizes
	rsaSignatureSize := len(rsaSignature)
	ecSignatureSize := len(ecR.Bytes()) + len(ecS.Bytes())
	
	fmt.Printf("Performance Comparison:\n")
	fmt.Printf("  RSA Signing Time: %v\n", rsaSignTime)
	fmt.Printf("  ECDSA Signing Time: %v\n", ecSignTime)
	fmt.Printf("  RSA Signature Size: %d bytes\n", rsaSignatureSize)
	fmt.Printf("  ECDSA Signature Size: %d bytes\n", ecSignatureSize)
	
	// Time verification
	startTime = time.Now()
	rsa.VerifyPSS(publicKey, crypto.SHA256, testHash[:], rsaSignature, nil)
	rsaVerifyTime := time.Since(startTime)
	
	startTime = time.Now()
	ecdsa.Verify(ecPublicKey, testHash[:], ecR, ecS)
	ecVerifyTime := time.Since(startTime)
	
	fmt.Printf("  RSA Verification Time: %v\n", rsaVerifyTime)
	fmt.Printf("  ECDSA Verification Time: %v\n", ecVerifyTime)
	
	fmt.Println("\n=== Document Signing Simulation ===")
	
	// Simulate signing a document with metadata
	document := struct {
		Content   string
		Author    string
		Timestamp string
		Version   int
	}{
		Content:   "Important contract terms and conditions...",
		Author:    "John Doe",
		Timestamp: "2024-01-15 10:30:00",
		Version:   1,
	}
	
	// Create document hash for signing
	documentData := fmt.Sprintf("Content:%s|Author:%s|Timestamp:%s|Version:%d",
		document.Content, document.Author, document.Timestamp, document.Version)
	documentHash := sha256.Sum256([]byte(documentData))
	
	fmt.Printf("Document to sign:\n")
	fmt.Printf("  Author: %s\n", document.Author)
	fmt.Printf("  Timestamp: %s\n", document.Timestamp)
	fmt.Printf("  Version: %d\n", document.Version)
	fmt.Printf("  Content: %s\n", document.Content)
	
	// Sign document with RSA
	docSignature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, documentHash[:], nil)
	if err != nil {
		fmt.Printf("Error signing document: %v\n", err)
		return
	}
	
	fmt.Printf("Document signed successfully\n")
	
	// Verify document signature
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, documentHash[:], docSignature, nil)
	if err != nil {
		fmt.Printf("❌ Document signature verification failed: %v\n", err)
	} else {
		fmt.Println("✅ Document signature verified - document is authentic")
	}
	
	fmt.Println("\n=== Key Management Best Practices ===")
	
	fmt.Println("Digital Signature Best Practices:")
	fmt.Println("✅ Always hash messages before signing")
	fmt.Println("✅ Use strong random number generation")
	fmt.Println("✅ Protect private keys with proper access controls")
	fmt.Println("✅ Use appropriate key sizes (RSA ≥ 2048 bits, ECDSA ≥ 256 bits)")
	fmt.Println("✅ Implement proper key rotation policies")
	fmt.Println("✅ Validate signatures before trusting signed data")
	fmt.Println("✅ Consider using established libraries and standards")
	
	fmt.Println("\nRSA vs ECDSA Comparison:")
	fmt.Println("RSA:")
	fmt.Println("  + Widely supported and standardized")
	fmt.Println("  + Simple key generation and usage")
	fmt.Println("  - Larger key sizes and signatures")
	fmt.Println("  - Slower performance")
	fmt.Println("ECDSA:")
	fmt.Println("  + Smaller keys and signatures")
	fmt.Println("  + Faster performance")
	fmt.Println("  + Better for mobile/embedded systems")
	fmt.Println("  - More complex implementation")
	fmt.Println("  - Requires secure random number generation")
}
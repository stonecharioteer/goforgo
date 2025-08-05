// digital_signatures.go
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
)

func main() {
	fmt.Println("=== Digital Signatures ===")
	
	fmt.Println("\n=== RSA Digital Signatures ===")
	
	// TODO: Generate RSA key pair
	privateKey, err := /* generate RSA private key (2048 bits) */
	if err != nil {
		fmt.Printf("Error generating RSA key: %v\n", err)
		return
	}
	
	publicKey := /* get public key from private key */
	
	fmt.Printf("Generated RSA key pair (size: %d bits)\n", privateKey.Size()*8)
	
	// TODO: Message to sign
	message := []byte("This is a message to be digitally signed")
	fmt.Printf("Original message: %s\n", message)
	
	// TODO: Create message hash
	hash := /* create SHA-256 hash of message */
	
	// TODO: Sign the hash using RSA private key
	signature, err := /* sign hash using RSA-PSS */
	if err != nil {
		fmt.Printf("Error signing message: %v\n", err)
		return
	}
	
	fmt.Printf("Signature created (length: %d bytes)\n", len(signature))
	
	// TODO: Verify the signature using RSA public key
	err = /* verify signature using RSA-PSS */
	if err != nil {
		fmt.Printf("❌ Signature verification failed: %v\n", err)
	} else {
		fmt.Println("✅ RSA signature verification successful")
	}
	
	// TODO: Test with tampered message
	tamperedMessage := []byte("This is a TAMPERED message to be digitally signed")
	tamperedHash := /* create SHA-256 hash of tampered message */
	
	err = /* verify original signature against tampered hash */
	if err != nil {
		fmt.Println("✅ Tampered message correctly rejected")
	} else {
		fmt.Println("❌ Tampered message incorrectly accepted")
	}
	
	fmt.Println("\n=== ECDSA Digital Signatures ===")
	
	// TODO: Generate ECDSA key pair using P-256 curve
	ecPrivateKey, err := /* generate ECDSA private key */
	if err != nil {
		fmt.Printf("Error generating ECDSA key: %v\n", err)
		return
	}
	
	ecPublicKey := /* get public key from private key */
	
	fmt.Printf("Generated ECDSA key pair (curve: %s)\n", ecPrivateKey.Curve.Params().Name)
	
	// TODO: Sign message with ECDSA
	ecMessage := []byte("ECDSA message for signing")
	ecHash := /* create SHA-256 hash of message */
	
	// TODO: Sign the hash using ECDSA
	r, s, err := /* sign hash using ECDSA */
	if err != nil {
		fmt.Printf("Error signing with ECDSA: %v\n", err)
		return
	}
	
	fmt.Printf("ECDSA signature created (r: %d bytes, s: %d bytes)\n", 
		len(r.Bytes()), len(s.Bytes()))
	
	// TODO: Verify ECDSA signature
	valid := /* verify ECDSA signature */
	if valid {
		fmt.Println("✅ ECDSA signature verification successful")
	} else {
		fmt.Println("❌ ECDSA signature verification failed")
	}
	
	// TODO: Test ECDSA with tampered message
	ecTamperedMessage := []byte("ECDSA TAMPERED message for signing")
	ecTamperedHash := /* create SHA-256 hash of tampered message */
	
	valid = /* verify original signature against tampered hash */
	if !valid {
		fmt.Println("✅ ECDSA tampered message correctly rejected")
	} else {
		fmt.Println("❌ ECDSA tampered message incorrectly accepted")
	}
	
	fmt.Println("\n=== Digital Signature Comparison ===")
	
	// TODO: Compare RSA vs ECDSA performance and signature sizes
	testMessage := []byte("Performance test message for signature comparison")
	testHash := /* create SHA-256 hash */
	
	// TODO: Time RSA signing
	startTime := /* get current time */
	rsaSignature, _ := /* sign with RSA */
	rsaSignTime := /* calculate elapsed time */
	
	// TODO: Time ECDSA signing  
	startTime = /* get current time */
	ecR, ecS, _ := /* sign with ECDSA */
	ecSignTime := /* calculate elapsed time */
	
	// TODO: Calculate signature sizes
	rsaSignatureSize := len(rsaSignature)
	ecSignatureSize := /* calculate combined size of r and s */
	
	fmt.Printf("Performance Comparison:\n")
	fmt.Printf("  RSA Signing Time: %v\n", rsaSignTime)
	fmt.Printf("  ECDSA Signing Time: %v\n", ecSignTime)
	fmt.Printf("  RSA Signature Size: %d bytes\n", rsaSignatureSize)
	fmt.Printf("  ECDSA Signature Size: %d bytes\n", ecSignatureSize)
	
	// TODO: Time verification
	startTime = /* get current time */
	/* verify RSA signature */
	rsaVerifyTime := /* calculate elapsed time */
	
	startTime = /* get current time */
	/* verify ECDSA signature */
	ecVerifyTime := /* calculate elapsed time */
	
	fmt.Printf("  RSA Verification Time: %v\n", rsaVerifyTime)
	fmt.Printf("  ECDSA Verification Time: %v\n", ecVerifyTime)
	
	fmt.Println("\n=== Document Signing Simulation ===")
	
	// TODO: Simulate signing a document with metadata
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
	
	// TODO: Create document hash for signing
	documentData := /* format document as string for hashing */
	documentHash := /* create SHA-256 hash */
	
	fmt.Printf("Document to sign:\n")
	fmt.Printf("  Author: %s\n", document.Author)
	fmt.Printf("  Timestamp: %s\n", document.Timestamp)
	fmt.Printf("  Version: %d\n", document.Version)
	fmt.Printf("  Content: %s\n", document.Content)
	
	// TODO: Sign document with RSA
	docSignature, err := /* sign document hash */
	if err != nil {
		fmt.Printf("Error signing document: %v\n", err)
		return
	}
	
	fmt.Printf("Document signed successfully\n")
	
	// TODO: Verify document signature
	err = /* verify document signature */
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
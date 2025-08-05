// base64_encoding.go - SOLUTION
// Learn Base64 encoding and decoding in Go

package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	fmt.Println("=== Base64 Encoding/Decoding ===")
	
	// Original data to encode
	originalText := "Hello, Go Base64 encoding!"
	originalBytes := []byte(originalText)
	
	fmt.Printf("Original text: %s\n", originalText)
	fmt.Printf("Original bytes: %v\n", originalBytes)
	
	// Standard Base64 encoding
	fmt.Println("\n=== Standard Base64 Encoding ===")
	
	encoded := base64.StdEncoding.EncodeToString(originalBytes)
	fmt.Printf("Encoded: %s\n", encoded)
	
	// Decode back to original
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatalf("Failed to decode: %v", err)
	}
	
	fmt.Printf("Decoded: %s\n", string(decoded))
	fmt.Printf("Decoded bytes: %v\n", decoded)
	
	// URL-safe Base64 encoding (no padding)
	fmt.Println("\n=== URL-Safe Base64 Encoding ===")
	
	urlEncoded := base64.URLEncoding.EncodeToString(originalBytes)
	fmt.Printf("URL encoded: %s\n", urlEncoded)
	
	urlDecoded, err := base64.URLEncoding.DecodeString(urlEncoded)
	if err != nil {
		log.Fatalf("Failed to decode URL encoding: %v", err)
	}
	
	fmt.Printf("URL decoded: %s\n", string(urlDecoded))
	
	// Raw Standard encoding (no padding)
	fmt.Println("\n=== Raw Standard Encoding (No Padding) ===")
	
	rawEncoded := base64.RawStdEncoding.EncodeToString(originalBytes)
	fmt.Printf("Raw encoded: %s\n", rawEncoded)
	
	rawDecoded, err := base64.RawStdEncoding.DecodeString(rawEncoded)
	if err != nil {
		log.Fatalf("Failed to decode raw encoding: %v", err)
	}
	
	fmt.Printf("Raw decoded: %s\n", string(rawDecoded))
	
	// Raw URL encoding (no padding)
	fmt.Println("\n=== Raw URL Encoding (No Padding) ===")
	
	rawURLEncoded := base64.RawURLEncoding.EncodeToString(originalBytes)
	fmt.Printf("Raw URL encoded: %s\n", rawURLEncoded)
	
	rawURLDecoded, err := base64.RawURLEncoding.DecodeString(rawURLEncoded)
	if err != nil {
		log.Fatalf("Failed to decode raw URL encoding: %v", err)
	}
	
	fmt.Printf("Raw URL decoded: %s\n", string(rawURLDecoded))
	
	// Working with binary data
	fmt.Println("\n=== Binary Data Encoding ===")
	
	binaryData := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x47, 0x6f, 0x21}
	fmt.Printf("Binary data: %v\n", binaryData)
	fmt.Printf("As string: %s\n", string(binaryData))
	
	binaryEncoded := base64.StdEncoding.EncodeToString(binaryData)
	fmt.Printf("Binary encoded: %s\n", binaryEncoded)
	
	binaryDecoded, err := base64.StdEncoding.DecodeString(binaryEncoded)
	if err != nil {
		log.Fatalf("Failed to decode binary: %v", err)
	}
	
	fmt.Printf("Binary decoded: %v\n", binaryDecoded)
	fmt.Printf("Binary decoded as string: %s\n", string(binaryDecoded))
	
	// Encoding/decoding with buffer
	fmt.Println("\n=== Buffer Operations ===")
	
	data := "This is a longer message that will be encoded using Base64 buffer operations!"
	fmt.Printf("Original: %s\n", data)
	
	// Calculate required buffer size for encoding
	encodedLen := base64.StdEncoding.EncodedLen(len(data))
	fmt.Printf("Required encoded buffer size: %d\n", encodedLen)
	
	// Encode to buffer
	encodedBuf := make([]byte, encodedLen)
	base64.StdEncoding.Encode(encodedBuf, []byte(data))
	fmt.Printf("Encoded in buffer: %s\n", string(encodedBuf))
	
	// Calculate required buffer size for decoding
	decodedLen := base64.StdEncoding.DecodedLen(len(encodedBuf))
	fmt.Printf("Maximum decoded buffer size: %d\n", decodedLen)
	
	// Decode from buffer
	decodedBuf := make([]byte, decodedLen)
	n, err := base64.StdEncoding.Decode(decodedBuf, encodedBuf)
	if err != nil {
		log.Fatalf("Failed to decode from buffer: %v", err)
	}
	
	fmt.Printf("Actual decoded length: %d\n", n)
	fmt.Printf("Decoded from buffer: %s\n", string(decodedBuf[:n]))
	
	// Common use cases
	fmt.Println("\n=== Common Use Cases ===")
	
	// 1. Encoding credentials for HTTP Basic Auth
	username := "admin"
	password := "secret123"
	credentials := username + ":" + password
	authHeader := base64.StdEncoding.EncodeToString([]byte(credentials))
	fmt.Printf("HTTP Basic Auth: %s\n", authHeader)
	
	// 2. Encoding small images or files for embedding
	fakeImageData := []byte("This would be binary image data...")
	imageEncoded := base64.StdEncoding.EncodeToString(fakeImageData)
	fmt.Printf("Image data encoded: %s\n", imageEncoded)
	
	// 3. URL-safe encoding for tokens
	token := "user123:session456:timestamp789"
	urlSafeToken := base64.URLEncoding.EncodeToString([]byte(token))
	fmt.Printf("URL-safe token: %s\n", urlSafeToken)
	
	// Verify all encodings decode back to original
	fmt.Println("\n=== Verification ===")
	
	testCases := []struct {
		name     string
		encoding *base64.Encoding
		encoded  string
	}{
		{"Standard", base64.StdEncoding, encoded},
		{"URL", base64.URLEncoding, urlEncoded},
		{"Raw Standard", base64.RawStdEncoding, rawEncoded},
		{"Raw URL", base64.RawURLEncoding, rawURLEncoded},
	}
	
	for _, tc := range testCases {
		decoded, err := tc.encoding.DecodeString(tc.encoded)
		if err != nil {
			fmt.Printf("❌ %s encoding failed: %v\n", tc.name, err)
		} else if string(decoded) == originalText {
			fmt.Printf("✓ %s encoding verified\n", tc.name)
		} else {
			fmt.Printf("❌ %s encoding mismatch\n", tc.name)
		}
	}
}
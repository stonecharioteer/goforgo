package main

import (
	"fmt"
	"time"
)

type User struct {
	ID   int
	Name string
}

func demonstrateRangeLoopTraps() {
	users := []User{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
	}
	
	fmt.Println("=== Goroutine Range Loop Trap - FIXED ===")
	// FIXED: Capture the loop variable in the goroutine
	for _, user := range users {
		// Method 1: Pass as parameter
		go func(u User) {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine processing user: %+v\n", u)
		}(user)
		
		// Method 2: Create local copy (alternative fix)
		// user := user // Shadow the loop variable
		// go func() {
		//     time.Sleep(100 * time.Millisecond)
		//     fmt.Printf("Goroutine processing user: %+v\n", user)
		// }()
	}
	time.Sleep(500 * time.Millisecond) // Wait for goroutines
	
	fmt.Println("\n=== Pointer Range Loop Trap - FIXED ===")
	var userPointers []*User
	
	// FIXED: Take address of a copy, not the loop variable
	for _, user := range users {
		userCopy := user // Create a copy
		userPointers = append(userPointers, &userCopy)
	}
	
	// Alternative fix: Use index instead of value
	var userPointers2 []*User
	for i := range users {
		userPointers2 = append(userPointers2, &users[i])
	}
	
	for i, ptr := range userPointers {
		fmt.Printf("User pointer %d: %+v\n", i, *ptr)
	}
	
	fmt.Println("\n=== Range Loop Modification - FIXED ===")
	numbers := []int{1, 2, 3, 4, 5}
	
	// CORRECT: Use index to modify original slice
	for i := range numbers {
		numbers[i] = numbers[i] * 2
	}
	fmt.Printf("Numbers after index modification: %v\n", numbers)
	
	// Alternative: Use index and value together
	numbers2 := []int{1, 2, 3, 4, 5}
	for i, num := range numbers2 {
		numbers2[i] = num * 2
	}
	fmt.Printf("Numbers2 after modification: %v\n", numbers2)
	
	fmt.Println("\n=== Map Range Order - Understanding the Behavior ===")
	userMap := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
		"David":   40,
		"Eve":     20,
	}
	
	fmt.Println("First iteration:")
	for name, age := range userMap {
		fmt.Printf("%s: %d\n", name, age)
	}
	
	fmt.Println("Second iteration:")
	for name, age := range userMap {
		fmt.Printf("%s: %d\n", name, age)
	}
	
	// UNDERSTANDING: Go randomizes map iteration order intentionally
	// If you need consistent order, collect keys and sort them:
	fmt.Println("Consistent order (keys collected and can be sorted):")
	var keys []string
	for k := range userMap {
		keys = append(keys, k)
	}
	// sort.Strings(keys) // Would give alphabetical order
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, userMap[k])
	}
}

func rangeWithChannelsFixed() {
	fmt.Println("\n=== Range with Channels - FIXED ===")
	
	ch := make(chan int, 5)
	
	// Send some values
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // IMPORTANT: Always close channel when done sending
	
	// Range over channel - works because channel is closed
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	
	// Demonstrate proper pattern with goroutine
	ch2 := make(chan int, 3)
	
	// Sender goroutine
	go func() {
		defer close(ch2) // FIXED: Always close in sender
		for i := 1; i <= 3; i++ {
			ch2 <- i
		}
	}()
	
	// Receiver can safely range
	for value := range ch2 {
		fmt.Printf("Received from ch2: %d\n", value)
	}
}

func stringRangeGotchaExplained() {
	fmt.Println("\n=== String Range Gotcha - EXPLAINED ===")
	
	// Ranging over strings iterates over runes, not bytes
	text := "Hello 世界" // Contains multi-byte UTF-8 characters
	
	fmt.Printf("String: %s\n", text)
	fmt.Printf("String length: %d bytes\n", len(text))
	fmt.Printf("String rune count: %d runes\n", len([]rune(text)))
	
	fmt.Println("Range over runes (what Go does by default):")
	for i, r := range text {
		fmt.Printf("Byte index %d: rune %c (Unicode: U+%04X)\n", i, r, r)
	}
	
	fmt.Println("Range over bytes (if you specifically need bytes):")
	for i := 0; i < len(text); i++ {
		fmt.Printf("Byte index %d: %d (hex: %02x)\n", i, text[i], text[i])
	}
	
	// UNDERSTANDING: The byte index jumps by 3 for the Chinese characters
	// because they are 3-byte UTF-8 sequences
}

func main() {
	demonstrateRangeLoopTraps()
	rangeWithChannelsFixed()
	stringRangeGotchaExplained()
	
	fmt.Println("\n=== All bugs fixed! Key takeaways: ===")
	fmt.Println("1. Always capture loop variables when using goroutines")
	fmt.Println("2. Don't take addresses of range loop variables")
	fmt.Println("3. Use indices to modify slice elements")
	fmt.Println("4. Map iteration order is randomized intentionally")
	fmt.Println("5. Always close channels when done sending")
	fmt.Println("6. String range iterates over runes, not bytes")
}
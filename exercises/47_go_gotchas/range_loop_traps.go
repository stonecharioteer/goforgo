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
	// Gotcha 1: Range loop variable reuse - the classic goroutine trap
	users := []User{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
	}
	
	fmt.Println("=== Goroutine Range Loop Trap ===")
	// This code has a bug - all goroutines will print the same user
	for _, user := range users {
		go func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine processing user: %+v\n", user)
		}()
	}
	time.Sleep(500 * time.Millisecond) // Wait for goroutines
	
	// Gotcha 2: Taking address of range loop variable
	fmt.Println("\n=== Pointer Range Loop Trap ===")
	var userPointers []*User
	
	// This code has a bug - all pointers will point to the same memory location
	for _, user := range users {
		userPointers = append(userPointers, &user)
	}
	
	for i, ptr := range userPointers {
		fmt.Printf("User pointer %d: %+v\n", i, *ptr)
	}
	
	// Gotcha 3: Modifying slice elements through range - doesn't work as expected
	fmt.Println("\n=== Range Loop Modification Trap ===")
	numbers := []int{1, 2, 3, 4, 5}
	
	// This loop tries to double all numbers but fails
	for _, num := range numbers {
		num = num * 2 // This doesn't modify the original slice
	}
	fmt.Printf("Numbers after range modification: %v\n", numbers)
	
	// This is the correct way (but still has the same bug pattern)
	for i, num := range numbers {
		numbers[i] = num * 2
	}
	fmt.Printf("Numbers after index modification: %v\n", numbers)
	
	// Gotcha 4: Range with map - iteration order is not guaranteed
	fmt.Println("\n=== Map Range Order Trap ===")
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
	// The order might be different!
}

func rangeWithChannels() {
	fmt.Println("\n=== Range with Channels ===")
	
	ch := make(chan int, 5)
	
	// Send some values
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // Important: must close channel for range to terminate
	
	// Range over channel
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	
	// What happens if we don't close the channel?
	ch2 := make(chan int, 3)
	ch2 <- 1
	ch2 <- 2
	// Missing close(ch2) - this would cause a deadlock if we ranged over it
	
	// Demonstrate the deadlock scenario (commented out to avoid hanging)
	// for value := range ch2 {
	//     fmt.Printf("This will deadlock: %d\n", value)
	// }
}

func stringRangeGotcha() {
	fmt.Println("\n=== String Range Gotcha ===")
	
	// Ranging over strings iterates over runes, not bytes
	text := "Hello 世界" // Contains multi-byte UTF-8 characters
	
	fmt.Printf("String length: %d bytes\n", len(text))
	
	// This iterates over runes (Unicode code points)
	for i, r := range text {
		fmt.Printf("Index %d: rune %c (Unicode: U+%04X)\n", i, r, r)
	}
	
	// If you want to iterate over bytes instead:
	for i := 0; i < len(text); i++ {
		fmt.Printf("Byte index %d: %d\n", i, text[i])
	}
	
	// The gotcha: index in range doesn't increment by 1 for multi-byte characters
}

func main() {
	demonstrateRangeLoopTraps()
	rangeWithChannels()
	stringRangeGotcha()
	
	fmt.Println("\n=== Fix the bugs in the above code! ===")
}
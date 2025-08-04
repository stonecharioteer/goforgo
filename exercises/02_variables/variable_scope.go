package main

import "fmt"

// TODO: Declare a package-level variable called globalMessage with value "Global"

func main() {
	// TODO: Declare a local variable called localMessage with value "Local"
	
	// This will create a new scope
	if true {
		// TODO: Declare a block-scoped variable called blockMessage with value "Block"
		fmt.Printf("%s %s %s\n", globalMessage, localMessage, blockMessage)
	}
	
	// blockMessage is not accessible here
	fmt.Printf("%s %s\n", globalMessage, localMessage)
}

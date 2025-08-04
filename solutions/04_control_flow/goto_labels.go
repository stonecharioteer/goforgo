package main

import "fmt"

func main() {
	// Use goto to skip some code
	// Print "Start"
	// Use goto to jump to label "skip"  
	// Print "This should be skipped"
	// Create label "skip:" and print "End"
	fmt.Println("Start")
	goto skip
	fmt.Println("This should be skipped")
skip:
	fmt.Println("End")
	
	// Use goto in a loop-like structure (not recommended, but educational)
	// Initialize counter := 0
	// Create label "loop:"
	// Print counter
	// Increment counter  
	// If counter < 3, goto loop
	// Print "Done with goto loop"
	fmt.Println("\nGoto loop:")
	counter := 0
loop:
	fmt.Println(counter)
	counter++
	if counter < 3 {
		goto loop
	}
	fmt.Println("Done with goto loop")
	
	// Use goto for error handling simulation
	// Create variables: success := false, attempt := 1
	// Create label "retry:"
	// Print "Attempt:", attempt
	// If attempt < 3 and !success:
	//   If attempt == 2, set success = true
	//   Increment attempt and goto retry
	// If success, print "Operation succeeded!"
	// Else print "Operation failed after 3 attempts"
	fmt.Println("\nError handling with goto:")
	success := false
	attempt := 1
retry:
	fmt.Println("Attempt:", attempt)
	if attempt < 3 && !success {
		if attempt == 2 {
			success = true
		}
		attempt++
		goto retry
	}
	if success {
		fmt.Println("Operation succeeded!")
	} else {
		fmt.Println("Operation failed after 3 attempts")
	}
	
	// Demonstrate goto with cleanup (common pattern)
	// Create variables: file := "data.txt", processed := false
	// Print "Opening file:", file
	// If file != "data.txt", goto cleanup
	// Print "Processing file..."
	// Set processed = true
	// Create label "cleanup:" 
	// Print "Cleaning up..."
	// If processed, print "File processed successfully"
	// Else print "File processing was skipped"
	fmt.Println("\nCleanup pattern with goto:")
	file := "data.txt"
	processed := false
	
	fmt.Println("Opening file:", file)
	if file != "data.txt" {
		goto cleanup
	}
	fmt.Println("Processing file...")
	processed = true
	
cleanup:
	fmt.Println("Cleaning up...")
	if processed {
		fmt.Println("File processed successfully")
	} else {
		fmt.Println("File processing was skipped")
	}
}
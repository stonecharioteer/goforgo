package main

import "fmt"

func main() {
	var name string = "Go"
	var version float64 = 1.21
	var isOpen bool = true
	var users int = 1000000
	
	fmt.Printf("Language: %s %.2f, Open: %t, Users: %d\n", name, version, isOpen, users)
}

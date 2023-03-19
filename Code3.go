package main

import "fmt"

// add returns the sum of x and y
func add(x, y int) int {
	return x + y
}

func main() {
	sum := add(42, 13)
	fmt.Printf("The sum of 42 and 13 is: %d", sum)
}

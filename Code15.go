package main

import "fmt"

const (
	// Define a huge number by shifting a 1 bit left 100 places.
	// This creates a binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100

	// Shift the huge number right 99 places to get a smaller number.
	// This results in a binary number that is 1 followed by 1 zero, or 2.
	Small = Big >> 99
)

func multiplyBy10AndAdd1(x int) int {
	return x*10 + 1
}

func multiplyBy0Point1(x float64) float64 {
	return x * 0.1
}

func main() {
	// Call the function multiplyBy10AndAdd1 with the constant Small.
	fmt.Println(multiplyBy10AndAdd1(Small))

	// Call the function multiplyBy0Point1 with the constant Small.
	fmt.Println(multiplyBy0Point1(float64(Small)))

	// Call the function multiplyBy0Point1 with the constant Big.
	fmt.Println(multiplyBy0Point1(float64(Big)))
}

package main

import "fmt"

func main() {
	numProblems := 7.0
	sqrtProblems := mySqrt(numProblems)
	fmt.Printf("Now you have %g problems.\n", sqrtProblems)
}

func mySqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - ((z*z - x) / (2 * z))
	}
	return z
}

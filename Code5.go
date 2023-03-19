package main

import "fmt"

func main() {
	var a, b string = "hello", "world"
	a, b = b, a
	fmt.Println(a, b)
}

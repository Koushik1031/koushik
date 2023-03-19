package main

import "fmt"

func split(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}

func main() {
	x, y := split(17)
	fmt.Println(x, y)
}

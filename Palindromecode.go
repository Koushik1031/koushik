package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "") // Remove spaces
	reversed := reverseString(s)
	return s == reversed
}

func reverseString(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for i, r := range s {
		runes[n-i-1] = r
	}
	return string(runes)
}

func main() {
	input := "A man a plan a canal Panama"
	if isPalindrome(input) {
		fmt.Printf("%q is a palindrome\n", input)
	} else {
		fmt.Printf("%q is not a palindrome\n", input)
	}
}

// Package calculator provides a library for simple calculations in Go.
package main1

import "fmt"

// Add takes two numbers and returns the result of adding them together.
func add(a, b int) int {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b int) int {
	return b - a
}

func main() {
	fmt.Println(add(2, 4))
}

package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Function with multiple returns
func calculations(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	multiply := n1 * n2
	return sum, multiply
}

func main() {
	fmt.Println("Sum: ", sum(10, 20))

	// Assigned function
	var f = func(txt string) string {
		fmt.Println("This is inside the assigned function")
		return txt + " appended"
	}

	fmt.Println(f("Text"))

	resultSum, resultMultiplication := calculations(2, 4)
	fmt.Println(resultSum, resultMultiplication)

}

package main

import "fmt"

func main() {

	// Variable assignments perform a copy of the value
	n1 := 1
	n2 := n1
	fmt.Println(n1, n2)
	n1++
	fmt.Println(n1, n2)

	// Pointer saves the memory address
	var n3 int
	var pointer *int

	n3 = 100
	pointer = &n3
	fmt.Println(n3, pointer)

	// Dereference pointer to access the value stored on the address
	n3 += 50
	fmt.Println(n3, *pointer)
}

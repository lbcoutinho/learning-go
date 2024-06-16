package main

import "fmt"

func main() {
	num := 15

	// Standard if clause
	if num > 15 {
		fmt.Println("Greater than 15")
	} else if num == 15 {
		fmt.Println("Number is 15")
	} else {
		fmt.Println("Less than 15")
	}

	// If init var (var limited to the if scope)
	if num2 := 5; num2 > 0 {
		fmt.Println("Number is positive")
	}
}

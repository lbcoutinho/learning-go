package main

import "fmt"

// Function with named returns
func calc(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}

// Variadic function
func printText(text string, numbers ...int) {
	for _, num := range numbers {
		fmt.Println(text, num)
	}
}

// Panic an recover
func recoverAvg() {
	if r := recover(); r != nil {
		fmt.Println("Function checkAvg was recovered")
	}
}

func checkAvg(n1, n2 int) {
	defer recoverAvg()

	avg := (n1 + n2) / 2
	if avg > 5 {
		fmt.Println("Avg greater than 5")
	} else if avg < 5 {
		fmt.Println("Avg less than 5")
	} else {
		panic("Avg is 5!!")
	}
}

// Closure function is bound to its internal variables
func closureSum() func(int) int {
	sum := 0
	return func(num int) int {
		sum += num
		return sum
	}
}

// Function with pointer as argument
func half(num *float32) {
	*num = *num / 2
}

func init() {
	fmt.Println("Init function executes before main")
}

func main() {
	// Defer will cause this function to be executed last
	defer fmt.Println(calc(10, 20))
	printText("Hello", 1, 2, 3, 4, 5)

	// Anonymous functions
	func(text string) {
		fmt.Println(text)
	}("Writing from anonymous function")

	checkAvg(2, 3)
	checkAvg(7, 8)
	checkAvg(5, 5)

	// Closure examples
	positiveSum, negativeSum := closureSum(), closureSum()
	fmt.Println("Positive sum")
	for i := 0; i < 5; i++ {
		fmt.Println(positiveSum(i))
	}
	fmt.Println("Negative sum")
	for i := 0; i > -5; i-- {
		fmt.Println(negativeSum(i))
	}

	// Sending memory address to function
	num := float32(5.0)
	fmt.Println(num)
	half(&num)
	fmt.Println(num)

}

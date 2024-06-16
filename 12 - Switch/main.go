package main

import "fmt"

func weekDay(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid"
	}
}

func printOperation(operation string, number int) {
	switch {
	case number == 0:
		fmt.Printf("Zero\n")
	case number < 0:
		fmt.Printf("Number is negative\n")
		fallthrough // Send execution directly inside to the case below
	case operation == "+":
		fmt.Printf("Sum %d\n", number)
	case operation == "-":
		fmt.Printf("Subtract %d\n", number)
	case operation == "*":
		fmt.Printf("Multiply %d\n", number)
	case operation == "/":
		fmt.Printf("Divide %d\n", number)
	}
}

func main() {

	fmt.Println(weekDay(3))

	printOperation("+", -10)
	printOperation("+", 3)

}

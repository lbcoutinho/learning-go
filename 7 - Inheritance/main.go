package main

import "fmt"

// THERE IS NO INHERITANCE IN GO...THIS IS JUST A WAY TO REUSE FIELDS ON ANOTHER STRUCTS
type person struct {
	name string
	age  uint8
}

type student struct {
	person
	college string
	course  string
}

func main() {
	person := person{"Steve", 15}

	student := student{person, "MIT", "Engineering"}

	fmt.Println(student)
}

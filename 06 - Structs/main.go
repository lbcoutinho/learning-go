package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
}

func main() {
	var u1 user
	fmt.Println(u1)

	u1.name = "Leandro"
	u1.age = 100
	fmt.Println(u1)

	// Inferred type
	address := address{"Rua 1"}
	u2 := user{"Foo", 20, address}
	fmt.Println(u2)

	// Init specific fields
	u3 := user{name: "Code"}
	fmt.Println(u3)

}

package main

import "fmt"

func main() {

	user := map[string]string{
		"name":    "Peter",
		"surname": "Silva",
	}

	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"first": "John",
			"last":  "Peter",
		},
	}

	fmt.Println(user2)

	// Delete a key from map
	fmt.Println(user)
	delete(user, "surname")
	fmt.Println(user)

	// Add key/value to map
	user["nickname"] = "P"
	fmt.Println(user)
}

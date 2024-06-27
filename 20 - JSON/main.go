package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint8  `json:"age"`
}

func marshal() {
	d := dog{"Toby", "Labrador", 8}

	text, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(text)
	fmt.Println(bytes.NewBuffer(text))
}

func unmarshal() {
	var d dog
	text := `{"name": "Beth", "breed": "Poodle", "age": 2}`

	if err := json.Unmarshal([]byte(text), &d); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)
}

func main() {
	marshal()
	unmarshal()
}

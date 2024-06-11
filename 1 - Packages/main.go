package main

import (
	"fmt"
	internal1 "module1/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing for main")
	internal1.Internal()

	error := checkmail.ValidateFormat("leandro@email.com")
	fmt.Println(error)
}

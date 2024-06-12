package main

import "fmt"

func main() {
	var variable1 string = "this is a var with explicit type"
	variable2 := "this is a var with implicit type"
	fmt.Println(variable1)
	fmt.Println(variable2)

	var (
		variable3 string = "var3 defined on a block"
		variable4 string = "var4 defined on a block"
	)

	fmt.Println(variable3, variable4)

	variable5, variable6 := "var5", "var6"
	fmt.Println(variable5, variable6)

	const constant1 = "this is constant"
	fmt.Println(constant1)

	variable5, variable6 = variable6, variable5
	fmt.Println("reversed values: ", variable5, variable6)
}

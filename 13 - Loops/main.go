package main

import (
	"fmt"
)

func main() {
	// i := 0

	// Like a while loop
	// for i < 5 {
	// 	i++
	// 	fmt.Println("Increment i:", i)
	// 	time.Sleep(time.Second)
	// }

	// // Standard for loop
	// for j := 0; j < 5; j++ {
	// 	fmt.Println("Increment j:", j)
	// 	time.Sleep(time.Second)
	// }

	// For using range
	names := [3]string{"Peter", "John", "Jim"}
	for index, name := range names {
		fmt.Println("Index ", index, " has name", name)
	}

	// Range on maps
	numbers := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	for key, value := range numbers {
		fmt.Println(key, value)
	}

	// Infinite loop
	for {
		fmt.Println("Stuck on infinity loop")
	}

}

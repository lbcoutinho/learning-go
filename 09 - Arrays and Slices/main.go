package main

import "fmt"

func main() {
	// Arrays - fixed length
	var arr1 [5]int
	arr1[0] = 1
	fmt.Println(arr1)

	arr2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3}
	fmt.Println(arr3)

	// Slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// append function appends the value to the end of the slide and return a new slice
	slice = append(slice, 6)
	fmt.Println(slice)

	// This is like a sublist function, start index is inclusive and end index is exclusive
	slice2 := arr2[1:3]
	fmt.Println(slice2)

	// Internal array
	slice3 := make([]int, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	slice3 = append(slice3, 11)
	slice3 = append(slice3, 12)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3)) // Go doubles the slice capacity once the limit is reached
}

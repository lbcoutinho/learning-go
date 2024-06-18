package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width  float32
	height float32
}

func (r rectangle) area() float32 {
	return r.height * r.width
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * (c.radius * c.radius)
}

type geometricShape interface {
	area() float32
}

func printArea(shape geometricShape) {
	fmt.Println("Area:", shape.area())
}

// Generic interface
func genericPrint(i interface{}) {
	fmt.Println(i)
}

func main() {
	rec := rectangle{2, 4}
	circ := circle{3}

	printArea(rec)
	printArea(circ)

	// Generic print accepts any type
	genericPrint(true)
	genericPrint("texto")
	genericPrint(1)
}

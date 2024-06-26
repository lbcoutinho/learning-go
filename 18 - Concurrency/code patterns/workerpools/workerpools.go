package main

import "fmt"

// Pattern worked pools can be used we you have a sequence of tasks to do then you can use the workers to performed the processing concurrently.

func fibonacci(pos int) int {
	if pos <= 1 {
		return pos
	}

	return fibonacci(pos-2) + fibonacci(pos-1)
}

func workers(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- fibonacci(number)
	}
}

func main() {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go workers(tasks, results)
	go workers(tasks, results)
	go workers(tasks, results)
	go workers(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		res := <-results
		fmt.Println(res)
	}
}

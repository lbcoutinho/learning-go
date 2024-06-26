package main

import (
	"fmt"
	"sync"
	"time"
)

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // Numbers of goroutines that are included on the wait group

	go func() {
		write("Hello!")
		waitGroup.Done() // Decrement wait group count by 1
	}()

	go func() {
		write("Go Programming!")
		waitGroup.Done()
	}()

	waitGroup.Wait() // Forces the application to wait for all the goroutines on the wait group to finish
	fmt.Println("Finished")
}

package main

import (
	"fmt"
	"time"
)

// Generator function is used to wrap a go routine execution an anonymous function. The generator returns a channel to output the function results.

func main() {
	channel := write("Hello!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Message is: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}

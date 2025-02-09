package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("Hello!", channel)

	// for {
	// 	msg, isOpen := <-channel
	// 	if !isOpen {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	for msg := range channel {
		fmt.Println(msg)
	}

	fmt.Println("Finished")
}

func write(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}

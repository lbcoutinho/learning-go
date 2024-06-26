package main

import (
	"fmt"
	"time"
)

func main() {
	chan1, chan2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			chan1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			chan2 <- "Channel 2"
		}
	}()

	for {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}
}

package examples

import (
	"fmt"
	"time"
)

func sendMessage(ch chan string, delay time.Duration, msg string) {
	time.Sleep(delay)
	ch <- msg
}
func Select() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go sendMessage(ch1, 1*time.Second, "Hello")
	go sendMessage(ch2, 1*time.Second, "World")
	for i := 1; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received :", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received :", msg2)
		}

	}
}

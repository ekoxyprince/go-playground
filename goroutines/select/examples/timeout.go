package examples

import (
	"fmt"
	"time"
)

func sendMsg(ch chan string, delay time.Duration, msg string) {
	time.Sleep(delay)
	ch <- msg
}
func Timeout() {
	ch1 := make(chan string)
	go sendMsg(ch1, 3*time.Second, "Trying out timeout with select statment")
	select {
	case msg1 := <-ch1:
		fmt.Println("Received :", msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout no response received")
	}

}

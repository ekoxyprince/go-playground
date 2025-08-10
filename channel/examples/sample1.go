package examples

import "fmt"

func SampleOne() {
	ch := make(chan string)
	go func() {
		ch <- "Hello from channels"
	}()
	msg := <-ch
	fmt.Println(msg)
}

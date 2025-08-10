package examples

import "fmt"

func BufferedChannel() {
	ch := make(chan string, 3)

	ch <- "Message1"
	ch <- "Message2"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

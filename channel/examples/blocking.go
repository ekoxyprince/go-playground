package examples

import "fmt"

func Blocking() {
	ch := make(chan string)
	go func() {
		// ch <- "Blocks go routine forever"
	}()
	fmt.Println("Should work forever")
	<-ch
	// fmt.Println(<-ch)
	fmt.Println("Will now work")
}

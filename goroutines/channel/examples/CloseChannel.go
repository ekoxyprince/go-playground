package examples

import "fmt"

func CloseChannel() {
	ch := make(chan int)

	go func() {
		for i := 0; i <= 3; i++ {
			ch <- 1
		}
		close(ch)
	}()
	for val := range ch {
		fmt.Println(val)
	}
}

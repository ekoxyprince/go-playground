package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0
var mu sync.Mutex

func increment() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}
func main() {
	go increment()
	go increment()
	time.Sleep(time.Second)
	fmt.Println("Final value is ", counter)
}

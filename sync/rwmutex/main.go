package main

import (
	"fmt"
	"sync"
	"time"
)

var price = 2000
var rwMx sync.RWMutex

func readPrice(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	rwMx.RLock() // Reader Lock
	fmt.Printf("Reader %d about to read price \n", id)
	time.Sleep(1 * time.Second)
	rwMx.RUnlock()
	fmt.Printf("Price is %d \n", price)
}
func writePrice(wg *sync.WaitGroup, id int, amount int) {
	defer wg.Done()
	fmt.Printf("Reader %d about to write price \n", id)
	rwMx.Lock()
	price += amount
	time.Sleep(3 * time.Second)
	rwMx.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go readPrice(&wg, i)
	}
	wg.Add(1)
	go writePrice(&wg, 2, 200)
	wg.Wait()
	fmt.Println("Current price is ", price)
	fmt.Println("All Operations are complete")
}

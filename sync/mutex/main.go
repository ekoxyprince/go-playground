package main

import (
	"fmt"
	"sync"
	"time"
)

var balance = 1000
var mu sync.Mutex

func withdrawAmount(amount int, wg *sync.WaitGroup) {
	//Execution of defer is always in the FORM (LIFO - Last In First Out)
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	if amount <= balance {
		time.Sleep(2 * time.Second)
		balance -= amount
		fmt.Println("Balace remaining", balance)
	} else {
		fmt.Println("Insufficient Current balance", balance)
	}

}
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go withdrawAmount(400, &wg)
	go withdrawAmount(300, &wg)
	go withdrawAmount(400, &wg)
	wg.Wait()
	fmt.Println("final value ", balance)
}

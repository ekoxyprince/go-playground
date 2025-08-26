package main

import (
	"fmt"
	"sync"
	"time"
)

func processOrder(orderId int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Order being processed", orderId)
	time.Sleep(2 * time.Second)
	fmt.Println("Order completed!")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go processOrder(i, &wg)
	}
	// !if called without wg.Add() it blocks the code execution indefinitely
	wg.Wait() //blocks until all order are complete
	fmt.Println("All orders processed")
}

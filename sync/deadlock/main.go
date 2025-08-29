package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m1, m2 sync.Mutex
	go func() {
		m1.Lock()
		fmt.Println("Go routine 1 is Locked m1")
		m2.Lock() //Waits forever is M2 is already locked
		fmt.Println("Go routine 1 is Locked m2")
		time.Sleep(2 * time.Second)
		m1.Unlock()
		m2.Unlock()
	}()
	go func() {
		m2.Lock()
		fmt.Println("Go routine 2 is Locked m2")
		m1.Lock() //Waits forever if M1 is already locked
		fmt.Println("Go routine 2 is Locked m1")
		time.Sleep(time.Second)
		m1.Unlock()
		m2.Unlock()
	}()
	time.Sleep(3 * time.Second)
	fmt.Scanln()
	fmt.Println("Completed.")
}

package examples

import (
	"fmt"
	"time"
)

func Worker(done chan bool) {
	fmt.Println("Starting Work....")
	time.Sleep(3 * time.Second)
	fmt.Println("Work completed...")
	done <- true
}

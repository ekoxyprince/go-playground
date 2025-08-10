package main

import (
	"fmt"

	"astrodev.com/channels/examples"
)

func main() {
	// done := make(chan bool)
	// go examples.Worker(done)
	// <-done
	examples.CloseChannel()
	fmt.Println("Everything completed")

}

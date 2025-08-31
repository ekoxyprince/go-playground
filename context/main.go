package main

import (
	"context"
	"fmt"
	"time"
)

func WithTimeout() {
	ctx, Cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer Cancel()
	select {
	case <-ctx.Done():
		fmt.Println("Timeout reached ", ctx.Err())
	}
}
func WithDeadline() {
	ctx, Cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer Cancel()
	select {
	case <-ctx.Done():
		fmt.Println("Dealine reached", ctx.Err())
	}
}

func main() {
	WithDeadline()
	WithTimeout()
}

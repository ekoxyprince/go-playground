package main

import (
	"testing"
)

func Add(num1, num2 int) int {
	return num1 + num2
}
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5 but go %d", result)
		return
	}
	t.Log("Test completed successfully")

}
func benchmarker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

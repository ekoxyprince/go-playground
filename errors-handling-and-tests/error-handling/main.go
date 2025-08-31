package main

import (
	"errors"
	"fmt"
	"os"
)

func readFile() (string, error) {
	file, err := os.ReadFile("file.txt")
	if err != nil {
		// ? wrapping an error using fmt.Errorf or errors.New
		// ! Note wrapping an error use %w so the new error created can be wrapped around the old one instead of %s.
		//return "", fmt.Errorf("Error ocuured %w", err)
		return "", err
	}
	return string(file), nil
}
func main() {
	file, err := readFile()
	if err != nil {
		//unwrapping error to return a response based on that error
		//? u can also use errors.As
		//! Note using errors.As can extract a particular error from the trees of error
		if errors.Is(err, os.ErrNotExist) {
			panic("File not found")
		}
		//! user recover to prevent desruptive actions of panic
		//recover()
		panic(err)
	}
	fmt.Println(file)
}

package main

import "fmt"

type speaker interface {
	speak()
}
type Cat struct{}

func (c Cat) speak() {
	fmt.Println("Meow...")
}

type Dog struct{}

func (d *Dog) speak() {
	fmt.Println("Woof...")
}

type Goat struct{}

func NewGoat() speaker {
	return &Goat{}
}
func (g *Goat) speak() {
	fmt.Println("Bleat...")
}
func makeNoise(s speaker) {
	s.speak()
}

//! Type assession using interface{} || any
func checkType(i interface{}) {
	//Here we do some time assession using .(Type)
	v, ok := i.(int)
	if ok {
		fmt.Println("It is an integer", v)
	} else {
		fmt.Println("Not an integer")
	}

}
func main() {
	d := Dog{}
	c := Cat{}
	g := NewGoat()
	makeNoise(&d)
	makeNoise(c)
	makeNoise(g)
	checkType(5)
	checkType(6.5)

}

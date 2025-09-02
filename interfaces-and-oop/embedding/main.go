package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) speak() {
	fmt.Println(a.Name, "Has spoken")
}

type Dog struct {
	Animal
}

func (d Dog) speak() {
	fmt.Println(d.Name, "Is speaking")
}

func main() {
	d := Dog{
		Animal: Animal{
			Name: "Dog",
		},
	}
	d.speak()
	d.Animal.speak()
	a := Animal{
		Name: "Goat",
	}
	a.speak()
}

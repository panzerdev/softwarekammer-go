package main

import "fmt"

type Animal struct {
	Kind string
}

// START OMIT
type Dog struct {
	Name   string
	Animal Animal
}

func (d Dog) Greet() {
	fmt.Printf("I am %v\n", d.Name)
}

func main() {
	laila := Dog{
		Name: "Laila",
		Animal: Animal{
			Kind: "Dog",
		},
	}
	laila.Greet()
}

// END OMIT

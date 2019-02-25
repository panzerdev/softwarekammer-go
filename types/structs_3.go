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

func (d Dog) Greet(s string) {
	fmt.Printf("I am %v\n", s)
}

func main() {
	laila := Dog{
		Name: "Laila",
		Animal: Animal{
			Kind: "Dog",
		},
	}
	laila.Greet(laila.Name)
}

// END OMIT

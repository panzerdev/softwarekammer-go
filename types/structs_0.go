package main

import "fmt"
// START OMIT
type Animal struct {
	Kind string
}
type Dog struct {
	Name   string
	Animal Animal
	Greet  func(string)
}

func main() {
	laila := Dog{
		Name: "Laila",
		Animal: Animal{
			Kind: "Dog",
		},
		Greet: func(s string) {
			fmt.Printf("I am %v\n", s)
		},
	}
	laila.Greet(laila.Name)
}
// END OMIT
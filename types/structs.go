package main

import "fmt"

type Animal struct {
	Kind string
}
type Dog struct {
	Name     string
	Age      int
	Birthday string
	Animal   Animal
}
func main() {
	laila := Dog{
		Name:     "Laila",
		Age:      13,
		Birthday: "01-10-2006",
		Animal: Animal{
			Kind: "Dog",
		},
	}
	fmt.Printf("%+v", laila)
}

package main

import (
	"errors"
	"fmt"
)

// START OMIT
type Dog struct{ name string }

func (d Dog) PrintName() {
	fmt.Println("The Dog is called:", d.name)
}
func (d *Dog) GiveName(n string) error {
	if d == nil {
		return errors.New("There is no Dog")
	}
	d.name = n
	return nil
}
func main() {
	var noDog *Dog
	err := noDog.GiveName("Laila")
	if err != nil {
		fmt.Println(err)
	}
	dog := Dog{}
	dog.GiveName("Bello")
	dog.PrintName()
	// END OMIT
}

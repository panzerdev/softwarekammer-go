package main

import "fmt"
//START OMIT
type Animal interface {
	Describe() string
}

type Doggy string
func (d Doggy) Describe() string {
	return string(d)
}

type Car string

func TellMeAboutAnimals(animal ...Animal) {
	for _, v := range animal {
		fmt.Println(v.Describe())
	}
}
func main() {
	dog := Doggy("Laila")
	car := Car("VW")
	TellMeAboutAnimals(dog, car)
}
// END OMIT
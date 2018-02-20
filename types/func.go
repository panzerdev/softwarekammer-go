package main

import (
	"fmt"
)

func Greet(name string) (bool, int, error) {
	nr, err := fmt.Println("Hello", name)
	return true, nr, err
}
func main() {
	fmt.Println(Greet("Michael 🤔"))
	f := Greet

	fmt.Println(f("Peter"))
	fmt.Println(MetaGreet("Hans", f))
}

func MetaGreet(name string, f func(string) (bool, int, error)) (bool, int, error) {
	b, i, err := f(name)
	return false, 0, fmt.Errorf("Always fail: %v, %v, %v", b, i, err)
}

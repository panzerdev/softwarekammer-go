package main

import "fmt"

func main() {
	var m map[string]int = map[string]int{}
	m["Peter"] = 12
	fmt.Println("Age:", m["Peter"])
	delete(m, "Peter")
	fmt.Println("Age:", m["Peter"])
	age, ok := m["Peter"]
	fmt.Println("Age:", age, "Valid", ok)
}

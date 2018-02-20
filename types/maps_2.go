package main

import "fmt"

func main() {
	m := map[string]int{"Peter": 12, "Michael": 34, "Leipzig": 999}
	fmt.Println("Size", len(m))
	for k := range m {
		fmt.Println("Key:", k)
	}
	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}
}

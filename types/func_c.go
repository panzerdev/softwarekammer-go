package main

import "fmt"

func main() {
	f := func(name string) {
		fmt.Println("Name:", name)
	}
	f("Robert")

	outer := "OuterValue"
	PassClosure("Bibo", func(s string) {
		fmt.Println(outer, "Closure", s)
	})
}
func PassClosure(name string, f func(string)) {
	f(name)
}

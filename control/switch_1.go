package main

import (
	"fmt"
)
func main() {
	TellType("Hello")
	TellType("Bye")
	TellType("No idea")
}
func TellType(v string) {
	switch v {
	case "Hello":
		fmt.Println("Greeting")
	case "Bye":
		fmt.Println("Godbye")
	default:
		fmt.Println("no idea!")
	}
}


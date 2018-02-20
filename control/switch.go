package main

import (
	"fmt"
	"time"
)
func main() {
	TellType(123)
	TellType("Hello World")
	TellType(time.Second)
	TellType(map[int]int{})
}
func TellType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("It's an int")
	case string:
		fmt.Println("It's a string")
	case time.Duration:
		fmt.Println("It's a druation")
	default:
		fmt.Println("no idea!")
	}
}

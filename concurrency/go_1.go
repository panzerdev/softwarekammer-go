package main

import (
	"fmt"
)

func main() {
	name := "Michael"
	go func() {
		fmt.Println("Say my name, say my name...", name)
	}()
}

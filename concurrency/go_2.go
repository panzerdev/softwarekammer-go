package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Michael"
	go func() {
		fmt.Println("Say my name, say my name...", name)
	}()

	time.Sleep(time.Second)
}

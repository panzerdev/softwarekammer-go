package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		go func() {
			fmt.Println("Your Nr:", i)
		}()
	}
	time.Sleep(time.Second)
}

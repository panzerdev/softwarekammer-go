package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go func(chan int) {
			time.Sleep(time.Second * time.Duration(rand.Intn(3)))
			c <- rand.Intn(10000)
		}(c)
	}

	result := 0
	for i := 0; i < 5; i++ {
		v := <-c
		fmt.Println("Read from chan", i, v)
		result += v
	}
	fmt.Println("Result sum", result)
}
// END OMIT

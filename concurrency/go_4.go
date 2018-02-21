package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// START OMIT
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(nr int) {
			time.Sleep(time.Second * time.Duration(rand.Intn(3)))
			fmt.Println("Your Nr:", nr)
			wg.Done()
		}(i)
	}
	fmt.Println("Start Waiting")
	wg.Wait()
	fmt.Println("Done")
}
// END OMIT
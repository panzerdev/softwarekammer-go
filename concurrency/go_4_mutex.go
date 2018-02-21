package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// START OMIT
func main() {
	mutex := sync.Mutex{}
	nrOfRoutines := 0
	for i := 0; i < 5; i++ {
		nrOfRoutines++
		go func(nr int) {
			time.Sleep(time.Second * time.Duration(rand.Intn(3)))
			fmt.Println("Your Nr:", nr)
			mutex.Lock()
			nrOfRoutines--
			mutex.Unlock()
		}(i)
	}
	fmt.Println("Start Waiting")
	for {
		mutex.Lock()
		if nrOfRoutines == 0 {
			break
		}
		mutex.Unlock()
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("Done")
}

// END OMIT

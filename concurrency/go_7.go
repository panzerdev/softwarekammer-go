package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	c := make(chan int)
	go StartCollecting(c)
	for {
		for i := 0; i < rand.Intn(1000); i++ {
			go emitData(c)
		}
		time.Sleep(time.Second)
	}
}

func StartCollecting(src <-chan int) {
	fmt.Println("Start collecting")
	data, tick := map[time.Time]int{}, time.NewTicker(time.Second*2)
	for {
		select {
		case v := <-src:
			data[time.Now()] = v
		case <-tick.C:
			fmt.Printf("Map has %v Items\n", len(data))
			data = map[time.Time]int{}
		}
	}
}

// END OMIT

func emitData(c chan<- int) {
	for i := 0; i < rand.Intn(1000); i++ {
		c <- i
	}
}

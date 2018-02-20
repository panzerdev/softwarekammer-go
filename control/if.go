package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	a := rand.Intn(15)
	if a <= 4 {
		fmt.Println(a, "Value < 5")
	} else if b := a + 1; b > 5 && b <= 10 {
		fmt.Println(a, "Value > 4 and < 10")
	} else {
		fmt.Println(a, "Value > 9")
	}
}

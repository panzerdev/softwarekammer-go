package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	res, err := GetGreatContentFast("http://some.url.here")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

// START OMIT
func GetGreatContentFast(url string) (string, error) {
	c := make(chan string, 3)
	for i := 0; i < cap(c); i++ {
		go get(url, c)
	}
	select {
	case res := <-c:
		return res, nil
	case <-time.After(time.Millisecond * 200):
		return "", fmt.Errorf("All Requests took too longer then: %v", time.Millisecond*200)
	}
}

// END OMIT

func get(url string, c chan<- string) {
	d := time.Duration(rand.Intn(1000)) * time.Millisecond
	fmt.Println(d)
	time.Sleep(d)
	c <- fmt.Sprintln(url, d.String())
}

package main

import "fmt"

func main() {
	nrs := []int{11, 22, 33, 44}
	for i := range nrs { // just index
		fmt.Println(i)
	}
	for i, v := range nrs { // index and value
		fmt.Println(i, v)
	}
}

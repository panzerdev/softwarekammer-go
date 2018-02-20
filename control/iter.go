package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	j := 0
	for j < 3 {
		fmt.Println(j)
		j++
	}
	for {
		fmt.Println(j)
		j++
		if j > 6 {
			break
		}
	}
}

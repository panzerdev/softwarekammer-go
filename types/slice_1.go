package main

import "fmt"

func main() {
	ids := []int{1, 2,}
	fmt.Println(ids, len(ids), cap(ids))

	ids = append(ids, 3, 4)
	fmt.Println(ids, len(ids), cap(ids))
}

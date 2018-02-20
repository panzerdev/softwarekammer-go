package main

import "fmt"

func main() {
	nr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nr, len(nr), cap(nr))
	// from including pos 4 to excluding pos 8
	slice1 := nr[4:8]
	slice2 := nr[:]  // from array to slice
	slice3 := nr[4:] // from 4 to end
	slice4 := nr[:8] // from start to 8
	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice3, len(slice3), cap(slice3))
	fmt.Println(slice4, len(slice4), cap(slice4))

	nrSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nrSlice, len(nrSlice), cap(nrSlice))
	nrSlice = append(nrSlice, 100)
	fmt.Println(nrSlice, len(nrSlice), cap(nrSlice))
}

package main

import "fmt"

func main() {
	nr := []int{1, 2, 3}
	fmt.Println("Before", nr)
	Change(nr)
	fmt.Println("After", nr)

	otherNr := nr
	otherNr[0], otherNr[1], otherNr[2] = 100, 200, 300
	fmt.Println(nr, otherNr)

	nrArray := [...]int{11, 22, 33}
	arrayCopy := nrArray
	arrayCopy[0], arrayCopy[1], arrayCopy[2] = 100, 200, 300
	fmt.Println(nrArray, arrayCopy)
}

func Change(a []int) {
	a[0], a[1], a[2] = 2, 4, 6
}

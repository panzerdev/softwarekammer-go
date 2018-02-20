package main

import "fmt"

func main() {
	nr := [3]int{1, 2, 3}
	fmt.Println(nr, len(nr), cap(nr))

	names := [...]string{"Michael", "Paul", "Jörg"}
	fmt.Println(names, len(nr), cap(nr))
}

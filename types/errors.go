package main

import (
	"fmt"
	"strconv"
)

func main() {
	nr, err := strconv.Atoi("HaHa")
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Println("Nr:", nr)
}

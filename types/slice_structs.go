package main

import "fmt"

type sliceItem struct{ i int }

func main() {
	a := []sliceItem{{i: 1}, {i: 2}}
	ChangeStruct(a)
	fmt.Println(a)
}

func ChangeStruct(v []sliceItem) {
	v[0].i = 123
	vc := v[1]
	vc.i = 222
}

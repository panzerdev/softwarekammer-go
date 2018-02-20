package main

import "fmt"

type sliceItemPointer struct{ i int }

func main() {
	a := []*sliceItemPointer{
		&sliceItemPointer{i: 1},
		&sliceItemPointer{i: 2},
	}
	ChangeStructPointer(a)
	fmt.Println(*a[0], *a[1])
}

func ChangeStructPointer(v []*sliceItemPointer) {
	v[0].i = 123
	vc := v[1]
	vc.i = 222
}

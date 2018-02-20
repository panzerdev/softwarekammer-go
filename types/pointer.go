package main

import "fmt"

type foo struct {
	bar string
}
func ByValue(v foo) {
	v.bar = "ByValue"
}
func ByReference(v *foo) {
	v.bar = "ByReference"
}
func main() {
	var v, r foo
	ByValue(v)
	ByReference(&r)
	fmt.Println("f", v.bar, "\nr", r.bar)
}

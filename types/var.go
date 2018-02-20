package main

var e = 1
var f int = 1

var (
	g int    = 1
	h string = "Hello"
)

var i string

func main() {
	var a int     // default value 0
	var b int = 1 // set to 1 and declare as 1
	var c = 1     // set to 1 and infer type from 1 (int)
	d := 1        // set to 1 and infer type from 1 (int)

	j, k, l := 1, false, "k is great" // set to 1 and infer type from 1 (int)
}

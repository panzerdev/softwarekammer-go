func main() {
	b, i err := Greet("First Greet")

	f := Greet
	b, i, err = f("My name")

	c := func(name string) (bool, int, error){
		nr, err := fmt.Println("Hello", name)
		return true, nr, err
	}
	b, i, err = c("Also my name")
}
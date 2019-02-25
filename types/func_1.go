func Greet(name string) (bool, int, error) {
	nr, err := fmt.Println("Hello", name)
	return true, nr, err
}
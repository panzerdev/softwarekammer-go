func MetaGreet(name string, f func(string) (bool, int, error)) (bool, int, error) {
	b, i, err := f(name)
	return b, i, err
}
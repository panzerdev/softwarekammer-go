func MetaGreet(name string, f func(string) (bool, int, error)) (bool, int, error) {
	return f(name)
}
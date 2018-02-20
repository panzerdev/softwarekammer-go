package dbh

// START OMIT
var world string

func InitWorld(w string) {
	setWorld(w)
}
func GetWorld() string {
	return world
}
func setWorld(w string) {
	world = w
}

// END OMIT

package main

import (
	"fmt"
	"github.com/panzerdev/softwarekammer-go/hello/dbh"
)

func main() {
	// START OMIT
	dbh.InitWorld("Earth")
	fmt.Println("Hello", dbh.GetWorld())
	// END OMIT
}

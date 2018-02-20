package main

// START OMIT
import (
	"fmt"
	"github.com/panzerdev/softwarekammer-go/hello/dbh"
)
// END OMIT

func main() {
	dbh.InitWorld("Earth")
	fmt.Println("Hello", dbh.GetWorld())
}

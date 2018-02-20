package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// START OMIT
	nr, err := strconv.Atoi("8080")
	if err != nil {
		return
	}
	_, err = http.Get(fmt.Sprintf("http://localhost:%v/ping", nr))
	if err != nil {
		return
	}

	nr, err = strconv.Atoi("8080")
	if err == nil {
		_, err = http.Get(fmt.Sprintf("http://localhost:%v/ping", nr))
		if err == nil {
			// more if err == nil to come...
		}
	}
	// END OMIT
}

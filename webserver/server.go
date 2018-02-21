package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// START OMIT
func main() {
	go func() {
		for {
			time.Sleep(time.Second)
			res, err := http.Get("http://localhost:8080/")
			if err != nil {
				continue
			}
			body, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			fmt.Println(string(body))
		}
	}()
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello Server World!")
	})
	http.ListenAndServe(":8080", nil)
}

// END OMIT

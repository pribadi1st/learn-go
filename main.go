package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Bad practice
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		log.Printf("Received: %s", err)
		if err != nil {
			http.Error(w, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello, %s", data)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Printf("Goodbye")
	})
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		GetImage(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

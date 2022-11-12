package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		Lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

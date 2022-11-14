package main

import (
	"fmt"
	"golang_book/chapter4/Tasks/xkcd"
	"log"
	"os"
)

func main() {
	comicses, err := xkcd.GetComics(os.Args[1:])
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, v := range comicses {
		fmt.Println(v.URL)
	}
}

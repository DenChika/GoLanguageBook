package main

import (
	"chapter4/Tasks/xkcd"
	"fmt"
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

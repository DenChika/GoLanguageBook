package main

import (
	"fmt"
	"golang_book/chapter4/Tasks/poster"
	"os"
	"strings"
)

func main() {
	title := strings.Join(os.Args[1:], " ")
	movie, err := poster.GetMovie(title)
	if err != nil {
		fmt.Printf("Сбой при получении фильма: %v", err)
	}
	poster.UploadPoster(movie)
}

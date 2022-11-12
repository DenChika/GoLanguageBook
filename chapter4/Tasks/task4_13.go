package main

import (
	"chapter4/Tasks/poster"
	"fmt"
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

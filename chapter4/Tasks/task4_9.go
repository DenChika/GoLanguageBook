package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordfreq() {
	input := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words[input.Text()]++
	}
	for key, value := range words {
		fmt.Printf("%s\t%d\n", key, value)
	}
}

func main() {
	wordfreq()
}

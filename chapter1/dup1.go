package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		count[input.Text()]++
	}

	for key, val := range count {
		if val > 1 {
			fmt.Printf("%d\t%s\n", val, key)
		}
	}
}

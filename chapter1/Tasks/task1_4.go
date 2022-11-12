package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for key, val := range counts {
		if val > 1 {
			fmt.Printf("%d\t%s\n", val, key)
		}
	}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "task1_4: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		flag := false
		for input.Scan() {
			if counts[input.Text()] > 1 {
				flag = true
				break
			}
		}
		if flag {
			fmt.Printf("%s\t", f.Name())
		}
		f.Close()
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

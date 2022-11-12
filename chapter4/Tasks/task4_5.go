package main

import "fmt"

func removeDuplicates(slice []string) []string {
	j := 0
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == slice[i+1] {
			slice[j] = slice[i+1]
		} else {
			slice[j] = slice[i]
			j++
		}
	}
	return slice[:(j + 1)]
}

func main() {
	data := []string{"hello", "hello", "hello", "lol", "lol", "kek", "kek", "lol", "kek", "kek", "kek"}
	data = removeDuplicates(data)
	fmt.Printf("%v\n", data)
}

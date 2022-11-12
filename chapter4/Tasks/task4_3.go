package main

import "fmt"

func reverseInt(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	reverseInt(&data)
	fmt.Printf("%v\n", data)
}

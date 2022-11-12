package main

import "fmt"

func main() {
	fmt.Println(IsAnagram("gglebik", "leggbik"))
}

func IsAnagram(s1 string, s2 string) bool {
	hash := make(map[byte]int)
	for _, ch := range s1 {
		hash[byte(ch)]++
	}
	for _, ch := range s2 {
		hash[byte(ch)]--
	}
	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}

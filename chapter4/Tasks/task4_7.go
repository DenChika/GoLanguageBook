package main

import (
	"fmt"
)

func reverseBytes(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	data := []byte{'a', 'b', 'o', 'b', 'u', 's'}
	reverseBytes(data)
	fmt.Printf("%v\n", data)
	//Без дополнительной памяти не обойтись, поскольку нам нужен bytes.Buffer
	//для преобразования в строку, а на вход поступает срез []byte.
	//Без буффера reverse выводит коды символов, а не сами символы
}

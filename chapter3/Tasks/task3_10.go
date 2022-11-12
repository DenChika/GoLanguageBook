package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma2("2345678910"))
}

func comma2(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for i, ch := range s {
		if (n-i)%3 == 0 && n > 3 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%c", ch)
	}
	return buf.String()
}

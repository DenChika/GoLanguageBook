package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma3("-2345678.910"))
}

func comma3(s string) string {
	var buf bytes.Buffer
	dot := strings.LastIndex(s, ".")
	sublen := dot
	for i, ch := range s {
		if sublen > 3 && i < dot && dot != -1 {
			if s[0] != '-' && (sublen-i)%3 == 0 || s[0] == '-' && (sublen-i)%3 == 0 && i != 0 {
				buf.WriteString(",")
			}
		}
		fmt.Fprintf(&buf, "%c", ch)
	}
	return buf.String()
}

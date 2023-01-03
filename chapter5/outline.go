package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for node := n.FirstChild; node != nil; node = n.NextSibling {
		outline(stack, node)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	outline(nil, doc)
}

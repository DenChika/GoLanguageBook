package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func nodeAmount(nodes map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		nodes[n.Data]++
	}

	for node := n.FirstChild; node != nil; node = n.NextSibling {
		nodeAmount(nodes, node)
	}
	return nodes
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "task5_2: %v\n", err)
	}
	fmt.Println(nodeAmount(nil, doc))
}

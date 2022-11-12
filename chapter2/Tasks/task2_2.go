package main

import (
	"bufio"
	"fmt"
	"golang_book/chapter2/tempconv"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		for _, arg := range strings.Split(input.Text(), " ") {
			convert(arg)
		}
	} else {
		for _, arg := range os.Args[1:] {
			convert(arg)
		}
	}
}

func convert(s string) {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
	}
	c := tempconv.Celsius(num)
	f := tempconv.Fahrenheit(num)
	k := tempconv.Kelvin(num)
	fmt.Printf("%s = %s = %s, %s = %s = %s, %s = %s = %s\n", c, tempconv.CToF(c), tempconv.CToK(c),
		f, tempconv.FToC(f), tempconv.FToK(f), k, tempconv.KToC(k), tempconv.KToF(k))
}

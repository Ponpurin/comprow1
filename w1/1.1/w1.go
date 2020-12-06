package main

import (
	"fmt"
	"strings"
)

func main() {
	var welc strings.Builder
	welc.WriteString("Hello")
	welc.WriteString(" ")
	welc.WriteString("World")
	fmt.Println(welc.String())
}

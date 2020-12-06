package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ReplaceAll("Hello World", "l", "x"))
	fmt.Println(strings.ReplaceAll("Hello World", "o", "z"))
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("Hello World", "l", "x", 2))
	fmt.Println(strings.Replace("Hello World", "l", "x", -1))
}

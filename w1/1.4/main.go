package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello World", "world"))
	fmt.Println(strings.Contains("Hello World", "World"))
}

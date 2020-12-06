package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("Hello World", "o"))
	fmt.Println(strings.Count("Hello World", ""))
}

package main

import "fmt"

func main() {
	var p *int
	i := 10
	fmt.Println("Value", i)
	p = &i
	*p = 3
	fmt.Println("Value", i)
}

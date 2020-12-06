package main

import "fmt"

func main() {
	numbers := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers[1])
	numbers[1] = 10
	fmt.Println(numbers[1])
	length := len(numbers)
	fmt.Println("Length = ", length)
}

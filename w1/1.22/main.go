package main

import "fmt"

func main() {
	alphabets := [2][3]string{{"A", "B", "C"}, {"a", "b", "c"}}
	fmt.Println(alphabets)
	fmt.Println(alphabets[0][1])

	numbers := [2][3][2]int{
		{
			{1, 2},
			{10, 20},
			{100, 200},
		},
		{
			{5, 6},
			{50, 60},
			{500, 600},
		},
	}
	fmt.Println(numbers)
	fmt.Println(numbers[0][2][1])
}

package main

import "fmt"

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sum(1, 2, 34, 44, 47, 5, 1224, 2345, 23, 534)
	fmt.Println(result)
}

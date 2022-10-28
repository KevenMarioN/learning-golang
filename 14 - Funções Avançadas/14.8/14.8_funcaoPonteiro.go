package main

import "fmt"

func inverteSinal(number *int) {
	*number = *number * -1
}

func main() {
	num := 20
	fmt.Println(num)
	inverteSinal(&num)

	fmt.Println(num)
}

package main

import "fmt"

func calcMaths(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	sum, sub := calcMaths(10, 25)
	fmt.Println(sum, sub)
}

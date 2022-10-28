package main

import (
	"fmt"
)

func sum(a int8, b int8) int8 {
	return a + b
}
func calcsMaths(n1, n2 int8) (int8, int8) {
	if n1 > n2 {
		return (n1 + n2), (n1 - n2)
	} else {
		return (n1 + n2), (n2 - n1)
	}
}

func main() {
	sumResult := sum(10, 20)
	fmt.Println(sumResult)

	var functionArrow = func(txt string) {
		fmt.Println(txt)
	}

	functionArrow("Keven Mário")

	resultSum, resultSubtraction := calcsMaths(7, 80)

	fmt.Println(resultSum)
	fmt.Println(resultSubtraction)
}

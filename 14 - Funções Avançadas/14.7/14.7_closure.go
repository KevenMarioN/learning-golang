package main

import "fmt"

func closure() func() {
	text := "Dentro da função clousure"

	function := func() {
		fmt.Println(text)
	}
	return function
}

func main() {
	text := "Dentro do func main"
	fmt.Println(text)

	closure()()
}

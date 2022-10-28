package main

import "fmt"

func fibonacci(posicion uint) uint {
	if posicion <= 1 {
		return posicion
	}

	return fibonacci(posicion-2) + fibonacci(posicion-1)
}

func main() {

	posicion := uint(6)

	fmt.Println(fibonacci(posicion))
}

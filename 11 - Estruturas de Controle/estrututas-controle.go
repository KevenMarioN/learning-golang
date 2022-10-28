package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controles")

	number := 10

	if number > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual que 15")
	}

	if outroNumber := number; outroNumber > 0 {
		fmt.Println("Número é maior que zero")
	}
}

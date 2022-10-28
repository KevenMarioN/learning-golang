package main

import "fmt"

func main() {
	result := func(text string) string {
		return fmt.Sprintf("Recebido => %v", text)
	}("Olá func anonima")

	fmt.Println(result)
}

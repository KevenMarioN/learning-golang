package main

import (
	"fmt"
	"test-introduction/addresses"
)

func main() {
	typeAddress := addresses.TypeAddress("Rua dos imigrantes")
	fmt.Println(typeAddress)
}

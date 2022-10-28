package main

import (
	"errors"
	"fmt"
)

func main() {
	var num int = 100000000
	fmt.Println(num)

	var num32 rune = 5896
	fmt.Println(num32)

	var numInt8 byte = 255
	fmt.Println(numInt8)

	numfloat := 580.78
	fmt.Println(numfloat)

	var str string = "Text"
	fmt.Println(str)

	var boolean bool
	fmt.Println(boolean)

	var erro error = errors.New("Internal Error")
	fmt.Println(erro)
}

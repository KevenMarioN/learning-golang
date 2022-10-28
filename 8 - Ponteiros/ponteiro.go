package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++

	fmt.Println(var1, var2)
	fmt.Println("Only var1 have alteration")

	var var3 int = 200
	var point *int

	point = &var3
	fmt.Println(var3, *point)
	*point = 23
	fmt.Println(var3)
}

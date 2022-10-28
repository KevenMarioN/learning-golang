package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	slice := []int{1, 5, 4, 5}

	fmt.Println(reflect.TypeOf((slice)))
	fmt.Println(reflect.TypeOf((array2)))

	fmt.Println(slice)
	slice = append(slice, 78)

	fmt.Println(slice)

	slice2 := array2[1:3]

	array2[1] = "Posição alterada"
	fmt.Println(slice2)

	// Array internos

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 32)
	slice3 = append(slice3, 22)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)

	fmt.Println(slice4)
	slice4 = append(slice4, 22)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}

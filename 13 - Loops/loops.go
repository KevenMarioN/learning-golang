package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 20 {
		// time.Sleep(time.Second)
		i++
	}
	fmt.Println(i)

	names := [3]string{"Keven", "Brendha", "Thales"}

	for _, name := range names {
		fmt.Println(name)
	}

	for _, letra := range "Keven" {
		fmt.Println(string(letra))
	}

	user := map[string]string{
		"name": "Keven",
		"age":  "22",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}
}

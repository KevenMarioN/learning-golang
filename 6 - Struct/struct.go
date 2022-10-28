package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number int
}

func main() {

	fmt.Println("Archive structs")
	var user1 user
	user1.age = 18
	user1.name = "Keven Mário"
	fmt.Println(user1)

	enderecoEx := address{street: "Rua dos Bobos"}

	user2 := user{
		name:    "Thales",
		age:     7,
		address: enderecoEx}

	fmt.Println(user2)

	user3 := user{age: 21}

	fmt.Println(user3)
}

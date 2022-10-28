package main

import "fmt"

type people struct {
	nome   string
	idade  uint8
	altura uint8
}

type student struct {
	people
	course  string
	college string
}

func main() {
	fmt.Println("Herança")

	p1 := people{"João", 18, 180}
	fmt.Println(p1)

	s1 := student{p1, "Engenharia", "USP"}
	fmt.Println(s1.nome)
}

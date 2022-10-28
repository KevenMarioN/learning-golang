package main

import "fmt"

type user struct {
	nome  string
	idade uint8
}

func (u user) save() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados \n", u.nome)
}

func (u *user) birth() {
	u.idade++
}

func main() {
	user1 := user{"Keven", 22}
	user1.save()
	user1.birth()
	fmt.Println(user1)
}

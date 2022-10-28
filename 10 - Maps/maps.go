package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{"nome": "Pedro"}

	fmt.Println(user)

	user2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"faculdade": "USP",
			"nome":      "engenharia",
		},
	}
	fmt.Println(user2)

	delete(user2, "nome")

	fmt.Println(user2)

	user2["hobbys"] = map[string]string{
		"dia-a-dia": "Surf",
		"erros":     "Teste",
	}

	fmt.Println(user2)
	delete(user2["hobbys"], "dia-a-dia")
	fmt.Println(user2)
}

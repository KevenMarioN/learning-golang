package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"-"`
}

func main() {
	dogJSON := []byte(`{"name" : "Rex","breed" : "Pit","age" : 2}`)

	var d dog

	fmt.Println(d)
	fmt.Println("Antes do Unmarshal")

	if err := json.Unmarshal(dogJSON, &d); err != nil {
		log.Fatal(err)
	}

	fmt.Println(d)
	fmt.Println("Depois do Unmarshal")
}

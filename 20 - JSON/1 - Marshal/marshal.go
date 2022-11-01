package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	d := dog{"Rex", "Pitbull", 2}
	fmt.Println(d)

	dJson, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dJson)

	fmt.Println(bytes.NewBuffer(dJson))

	d2 := map[string]string{
		"name":  "Huck",
		"breed": "Dalmata",
		"age":   "4",
	}

	d2Json, err := json.Marshal(d2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d2Json)
	fmt.Println(bytes.NewBuffer(d2Json))
}

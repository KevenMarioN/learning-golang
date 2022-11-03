package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("*.html"))

type user struct {
	Name  string
	Email string
}

func Home(write http.ResponseWriter, request *http.Request) {
	user := user{"Keven Mário", "kevenmario.n.r@gmail.com"}
	templates.ExecuteTemplate(write, "home.html", user)
}

func main() {
	http.HandleFunc("/home", Home)
	fmt.Println("Listinig in port 5000 🔥!")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

package main

import (
	"log"
	"net/http"
)

func Home(write http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		write.WriteHeader(404)
		write.Write([]byte("Method not Exists"))
	} else {
		write.Write([]byte("Method GET"))
	}

}

func main() {

	http.HandleFunc("/home", Home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

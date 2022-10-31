package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Olá mundo!"
	channel <- "Burfer"
	// channel <- "Olá mundo!"

	menssage := <-channel
	fmt.Println(menssage)

}

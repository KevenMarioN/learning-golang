package main

import (
	"fmt"
	"time"
)

func main() {
	menssage := write("Keven Mário")
	for i := 0; i < 10; i++ {
		fmt.Println(<-menssage)
	}
}

func write(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("O valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}

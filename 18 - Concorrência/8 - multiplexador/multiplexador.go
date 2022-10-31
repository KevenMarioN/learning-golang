package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	result := multiplex(write("Chan 1"), write("Chan 2"))
	for i := 0; i < 15; i++ {
		fmt.Println(<-result)
	}
}

func multiplex(channelInput1, channelInput2 <-chan string) <-chan string {
	channelOutput := make(chan string)
	go func() {
		for {
			select {
			case menssage := <-channelInput1:
				channelOutput <- menssage
			case menssage := <-channelInput2:
				channelOutput <- menssage
			}
		}
	}()
	return channelOutput
}

func write(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("O valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return channel
}

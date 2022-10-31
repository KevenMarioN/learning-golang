package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go write("Keven Mário", canal)

	for menssage := range canal {
		fmt.Println(menssage)
	}
}

func write(text string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- text
		time.Sleep(time.Second)
	}

	close(canal)
}

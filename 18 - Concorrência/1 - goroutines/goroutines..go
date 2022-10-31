package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Ola Mundo !")
	write("Programando em GO!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

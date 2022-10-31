package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		write("Ola Mundo !")
		waitGroup.Done()
	}()

	go func() {
		write("Programando em GO!")
		waitGroup.Done()
	}()

	go func() {
		write("GoRoutine 3!")
		waitGroup.Done()
	}()

	go func() {
		write("GoRoutine 4!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text, i)
		time.Sleep(time.Second)
	}
}

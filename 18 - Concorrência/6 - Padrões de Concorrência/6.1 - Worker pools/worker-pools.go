package main

import "fmt"

func main() {
	const number int = 10
	tasks, results := make(chan int, number), make(chan int, number)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < number; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < number; i++ {
		result := <-results
		fmt.Println(result)
	}
}

func worker(taks <-chan int, results chan<- int) {
	for number := range taks {
		results <- fibonacci(number)
	}
}
func fibonacci(posicion int) int {
	if posicion <= 1 {
		return posicion
	}

	return fibonacci(posicion-2) + fibonacci(posicion-1)
}

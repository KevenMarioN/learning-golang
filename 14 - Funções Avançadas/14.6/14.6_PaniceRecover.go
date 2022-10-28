package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func studentApproved(n1, n2 float64) bool {
	defer recoverExecution()
	medium := (n1 + n2) / 2

	if medium < 6 {
		return true
	} else if medium < 6 {
		return false
	}
	panic("A MÈDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(studentApproved(6, 6))
	fmt.Println("Pós execução!")
}

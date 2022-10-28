package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("MARIO")
	generic(111)
	generic(true)
}

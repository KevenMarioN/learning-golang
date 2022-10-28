package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	app := app.Construction()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

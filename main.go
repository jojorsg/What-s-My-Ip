package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicação := app.Gerar()
	if erro := aplicação.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
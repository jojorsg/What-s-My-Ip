package app

import "github.com/urfave/cli"

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "What's My IP"
	app.Usage = "Busca IP e servidor host"

	return app
}
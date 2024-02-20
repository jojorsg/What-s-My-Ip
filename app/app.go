package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "What's My IP"
	app.Usage = "Busca IP e servidor host"

	flags := []cli.Flag {
		cli.StringFlag {
			Name:"host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Busca IP de endereço web",
			Flags: flags,
			Action: buscarIp,
		},
		{
			Name: "servidores",
			Usage: "Busca o servidor host",
			Flags: flags,
			Action: buscarServidor,
		},
	}
	return app
}

func buscarIp(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidor(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
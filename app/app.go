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

	app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Busca IP de endereço web",
			Flags: []cli.Flag {
				cli.StringFlag {
					Name:"host",
					Value: "google.com",
				},
			},
			Action: buscarIp,
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
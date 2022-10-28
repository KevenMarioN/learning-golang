package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Construction returning the application the command line for execution
func Construction() *cli.App {
	app := cli.NewApp()

	app.Name = "Application the command line"
	app.Usage = "Search the ip's the websites"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:     "host",
			Required: true,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPS de address the internet",
			Flags:  flags,
			Action: searchIPS,
		},
		{
			Name:   "server",
			Usage:  "Search the name servers the internet",
			Flags:  flags,
			Action: searchServer,
		},
	}
	return app
}

func searchIPS(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
func searchServer(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, serve := range servers {
		fmt.Println(serve.Host)
	}

}

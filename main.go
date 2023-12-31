package main

import (
	"fmt"
	"os"
	"log"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "Healthchecker",
		Usage: "A tiny tool that checks whether a website is running or is down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "domain",
				Aliases: []string{"d"},
				Usage: "Domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name: "port",
				Aliases: []string{"p"},
				Usage: "port number to check.",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error{
			port:=c.String("port")
			if c.String("port")==""{
				port="80"
			}
			status:=Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err:=app.Run(os.Args)
	if err!=nil{
		log.Fatal(err)
	}
}
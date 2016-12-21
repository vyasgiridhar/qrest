package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/vyasgiridhar/qrest"
	"github.com/vyasgiridhar/qrest/config"
)

func main() {
	app := cli.NewApp()
	app.Name = "qrest"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Value: 8000,
			Usage: "Port to listen.",
		},
		cli.StringFlag{
			Name:  "host, h",
			Value: "127.0.0.1",
			Usage: "Maria db url",
		},
		cli.StringFlag{
			Name:  "user, u",
			Value: "",
			Usage: "Maria DB url",
		},
		cli.StringFlag{
			Name:  "pass, p",
			Value: "",
			Usage: "Maria DB url",
		},
		cli.IntFlag{
			Name:  "mport, mp",
			Value: 3306,
			Usage: "Maria DB port",
		},
		cli.StringFlag{
			Name:  "database, db",
			Value: "",
			Usage: "Maria DB Name",
		},
	}
	app.Action = func(c *cli.Context) {
		args := c.Args()
		if c.Int("port") == 0 || c.String("host") == "" || c.Int("mport") == 0 || c.String("user") == "" || c.String("pass") == "" || c.String("database") == "" {
			fmt.Println("All arguments are required")
			return
		}
		if args.Present() {
			qrest.Start(config.Config{
				c.Int("port"),
				c.String("host"),
				c.Int("mport"),
				c.String("user"),
				c.String("pass"),
				c.String("database"),
			})
		}
	}
	app.Run(os.Args())

}

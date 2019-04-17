package main

import (
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"go-stmp/pkg/utils"
	"go-stmp/pkg/stmp"
    "log"
	"os"
	"fmt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := NewApp()
	err = app.Run(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "go-stmp"
	app.Usage = "A ugly go send mail tool"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "env, e",
			Value: "STMP_HOST",
			Usage: "specify an environment variable containing the stmp host",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Generate a new mail template file",
			Action:  func(c *cli.Context) error { 
				name := c.Args().First()
				utils.NewMailTemplate(name)
				return nil 
			},
		},
		{
			Name:    "send",
			Aliases: []string{"s"},
			Usage:   "Send mail",
			Action:  func(c *cli.Context) error { 
				name := c.Args().First()
				stmp.SendMail(name)
				return nil 
			},
		},
	}

	return app
	
}
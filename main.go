package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	state, err := getLaboState()
	if err != nil {
		panic(err)
	}

	db, err := newDatabase(state.DatabaseFile)
	if err != nil {
		panic(err)
	}

	labo := &Labo{
		Version:  "0.0.2",
		Database: db,
	}

	app := cli.NewApp()
	app.Name = "labo"
	app.Version = "0.0.2"
	app.Usage = "create a simple folder for your simple projects"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "author, a",
			Value: "",
			Usage: "author of your project (probably your name)",
		},
		cli.StringFlag{
			Name:  "project-version, pv",
			Value: "0.0.1",
			Usage: "version of your project",
		},
		cli.StringFlag{
			Name:  "seed, s",
			Value: "https://github.com/bregydoc/labo-agnostic",
			Usage: "set your seed to scaffold your project",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create your new project",
			Action: func(c *cli.Context) error {
				name := c.Args().Get(0)

				seedURL := c.GlobalString("seed")
				if _, err := labo.CreateNewProject(name, seedURL); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "list all your projects",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

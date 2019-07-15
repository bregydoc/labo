package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Name = "labo"
	app.Version = "0.1"
	app.Usage = "create a simple folder for your simple projects"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "author, a",
			Value: hostname,
			Usage: "author of your project (probably your name)",
		},
		cli.StringFlag{
			Name:  "project-version, pv",
			Value: "0.0.1",
			Usage: "version of your project",
		},
	}

	app.Action = func(c *cli.Context) error {
		author := c.String("author")
		version := c.String("version")
		projectName := c.Args().Get(0)
		project := &Project{
			Name:      projectName,
			Age:       time.Now().Year(),
			Author:    author,
			CreatedAt: time.Now(),
			Version:   version,
		}

		fmt.Printf("Creating project '%s' by %s\n", project.Name, project.Author)
		if err := inflateProject(project); err != nil {
			return err
		}
		fmt.Printf("All done!\n")
		fmt.Printf("To go to your project, type:\n")
		fmt.Printf("  $ cd %s\n", project.Name)
		return nil
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"

	au "github.com/logrusorgru/aurora"
)

func (l *Labo) renderCreatingProject(project *Project) error {
	fmt.Printf("Creating project '%s' by %s\n", project.Name, project.Author)
	return nil
}

func (l *Labo) renderProjectCreated(project *Project) error {
	fmt.Printf("All done!\n")
	fmt.Printf("To go to your project, type:\n")
	fmt.Println(au.BrightGreen(fmt.Sprintf("  $ cd %s\n", project.Name)))
	return nil
}

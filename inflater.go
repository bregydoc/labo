package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	_ "github.com/bregydoc/labo/data"
	"github.com/phogolabs/parcello"
)

func createLicense(project *Project) error {
	f, err := parcello.Open("basic/LICENSE")
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	t, err := template.New("license").Parse(string(data))
	if err != nil {
		return err
	}
	buff := bytes.NewBufferString("")
	err = t.Execute(buff, project)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(project.ProjectName, "LICENSE"), buff.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func createReadme(project *Project) error {
	f, err := parcello.Open("basic/README.md")
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	t, err := template.New("readme").Parse(string(data))
	if err != nil {
		return err
	}
	buff := bytes.NewBufferString("")
	err = t.Execute(buff, project)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(project.ProjectName, "README.md"), buff.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func createGitignore(project *Project) error {
	f, err := parcello.Open("basic/.gitignore")
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	t, err := template.New("readme").Parse(string(data))
	if err != nil {
		return err
	}
	buff := bytes.NewBufferString("")
	err = t.Execute(buff, project)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(project.ProjectName, ".gitignore"), buff.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func inflateProject(project *Project) error {
	if project.ProjectName == "" {
		return fmt.Errorf("invalid project name")
	}
	err := os.Mkdir(project.ProjectName, os.ModePerm)
	if err != nil {
		return err
	}

	err = createLicense(project)
	if err != nil {
		return err
	}

	err = createReadme(project)
	if err != nil {
		return err
	}

	err = createGitignore(project)
	if err != nil {
		return err
	}

	return nil
}

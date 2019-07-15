package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/yaml.v2"
)

// name: agnostic
// gen:
//   - LICENSE
//   - README.md
// dirs:

// Seed represents a new seed (a.k.a. template) of labo
type Seed struct {
	url         string
	source      string
	initialized bool
	Name        string   `yaml:"name"`
	Gen         []string `yaml:"gen"`
	Dirs        []string `yaml:"dirs"`
}

func newSeed(url string) *Seed {
	return &Seed{url: url}
}

func (s *Seed) fetch() error {
	dir, err := ioutil.TempDir(os.TempDir(), "seed")
	if err != nil {
		return err
	}

	if s.url == "" {
		return errors.New("invalid url")
	}

	// absDir := path.Join(dir, s.Name)

	chunks := strings.Split(s.url, "/")
	repoName := chunks[len(chunks)-1]
	repoDir := path.Join(dir, repoName)

	_, err = git.PlainClone(repoDir, false, &git.CloneOptions{
		URL:      s.url,
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(path.Join(repoDir, ".labo.yml"))
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, s)
	if err != nil {
		return err
	}

	s.source = repoDir
	s.initialized = true

	return nil
}

func (s *Seed) gen(where string, project *Project) error {
	for _, g := range s.Gen {

		file, err := os.OpenFile(path.Join(where, g), os.O_RDONLY, 0644)
		if err != nil {
			return err
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		if err = file.Close(); err != nil {
			return err
		}

		file, err = os.OpenFile(path.Join(where, g), os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		temp, err := template.New(g).Parse(string(data))
		if err != nil {
			return err
		}

		if err = temp.Execute(file, project); err != nil {
			return err
		}

		if err = file.Close(); err != nil {
			return err
		}

	}
	return nil
}

func (s *Seed) dirs(where string) error {
	for _, dir := range s.Dirs {
		if err := os.Mkdir(path.Join(where, dir), 0755); err != nil {
			return err
		}
	}
	return nil
}

func (s *Seed) plant(where string, project *Project) error {
	if !s.initialized {
		return errors.New("seed not initialized")
	}
	var err error
	name := strings.ToLower(project.Name)

	destination := path.Join(where, name)

	if err = os.Mkdir(destination, 0755); err != nil {
		if os.IsExist(err) {
			fmt.Println("Your project already exist")
			return nil
		}
		return err
	}

	if err = copyDirectory(s.source, destination); err != nil {
		return err
	}

	if _, err = os.Open(path.Join(destination, ".git")); err == nil {
		if err = os.RemoveAll(path.Join(destination, ".git")); err != nil {
			return err
		}
	}

	if err = os.RemoveAll(path.Join(destination, ".labo.yml")); err != nil {
		return err
	}

	if err = s.gen(destination, project); err != nil {
		return err
	}

	if err = s.dirs(destination); err != nil {
		return err
	}

	project.Seed = s

	return nil
}

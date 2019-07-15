package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	"gopkg.in/src-d/go-git.v4"
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

	absDir := path.Join(dir, s.Name)

	_, err = git.PlainClone(absDir, false, &git.CloneOptions{
		URL:      s.url,
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	chunks := strings.Split(s.url, "/")
	repoName := chunks[len(chunks)-1]

	repo := path.Join(absDir, repoName)

	s.source = repo
	s.initialized = true

	return nil

}

func (s *Seed) gen(project *Project) error {
	for _, g := range s.Gen {
		file, err := os.OpenFile(path.Join(s.source, g), os.O_RDWR, 0644)
		if err != nil {
			return err
		}

		if err = template.New("template").Execute(file, project); err != nil {
			return err
		}

		if err = file.Close(); err != nil {
			return err
		}

	}
	return nil
}

func (s *Seed) dirs() error {
	for _, dir := range s.Dirs {
		if err := os.Mkdir(dir, 0466); err != nil {
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

	name := strings.ToLower(s.Name)

	destination := path.Join(where, name)

	if err = copyDirectory(s.source, destination); err != nil {
		return err
	}

	if _, err = os.Open(path.Join(destination, ".git")); err == nil {
		if err = os.RemoveAll(path.Join(destination, ".git")); err != nil {
			return err
		}
	}

	if err = s.gen(project); err != nil {
		return err
	}

	if err = s.dirs(); err != nil {
		return err
	}

	project.Seed = s

	return nil
}

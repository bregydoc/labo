package main

import (
	"time"
)

// Labo is the main program thread
type Labo struct {
	Version  string
	Database Database
}

// ProjectOption represents a project modifier
type ProjectOption interface {
	Type() string
	Value() interface{}
}

func (l *Labo) newProject(name, where, seedURL string, options ...ProjectOption) (*Project, error) {
	project, err := newDefault(name)
	if err != nil {
		return nil, err
	}

	if err = l.renderCreatingProject(project); err != nil {
		return nil, err
	}

	for _, opt := range options {
		switch opt.Type() {
		case "version":
			val, _ := opt.Value().(string)
			project.Version = val
		case "author":
			val, _ := opt.Value().(string)
			project.Author = val
		case "icon":
			val, _ := opt.Value().(string)
			project.Icon = val
		}
	}

	seed := newSeed(seedURL)

	if err = seed.fetch(); err != nil {
		return nil, err
	}

	if err = project.plant(where, seed); err != nil {
		return nil, err
	}

	project.State = Active

	project.UpdatedAt = time.Now()

	if err = l.Database.RegisterProject(project); err != nil {
		return nil, err
	}

	return project, nil
}

// CreateNewProject scaffold a new project, the only requeried param is the name of project
func (l *Labo) CreateNewProject(name, seed string, options ...ProjectOption) (*Project, error) {
	pwd, err := getPwd()
	if err != nil {
		return nil, err
	}

	project, err := l.newProject(name, pwd, seed, options...)
	if err != nil {
		return nil, err
	}

	if err = l.renderProjectCreated(project); err != nil {
		return nil, err
	}

	return project, nil

}

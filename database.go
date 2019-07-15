package main

type Database interface {
	RegisterProject(p *Project) error
	GetProjectByID(id string) (*Project, error)
	GetProjectsByName(name string) ([]*Project, error)
	GetAllProjects(states ...ProjectState) ([]*Project, error)
}

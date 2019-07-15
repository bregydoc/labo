package main

type Database interface {
	RegisterProject(p *Project) error
	GetProjectByID(id string) (*Project, error)
	GetProjectsByName(name string) ([]*Project, error)
	GetAllProjects(states ...ProjectState) ([]*Project, error)
	UpdateProject(id string, update Project) (*Project, error)
	DeleteProject(id string) (*Project, error)
}

package main

type Database interface {
	RegisterProject(p *Project) error
}

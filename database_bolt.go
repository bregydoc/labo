package main

import (
	"github.com/asdine/storm"
)

type boltDB struct {
	db   *storm.DB
	path string
}

func newDatabase(path string) (*boltDB, error) {
	db, err := storm.Open(path)
	if err != nil {
		return nil, err
	}

	return &boltDB{
		path: path,
		db:   db,
	}, nil
}

func (bolt *boltDB) RegisterProject(p *Project) error {
	return bolt.db.Save(p)
}

func (bolt *boltDB) GetProjectByID(id string) (*Project, error) {
	project := new(Project)
	if err := bolt.db.One("ID", id, project); err != nil {
		return nil, err
	}
	return project, nil
}

func (bolt *boltDB) GetProjectsByName(name string) ([]*Project, error) {
	projects := make([]*Project, 0)
	if err := bolt.db.Find("Name", name, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (bolt *boltDB) GetAllProjects(states ...ProjectState) ([]*Project, error) {
	projects := make([]*Project, 0)
	if err := bolt.db.All(&projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (bolt *boltDB) UpdateProject(id string, update Project) (*Project, error) {
	panic("unimplemented")
}

func (bolt *boltDB) DeleteProject(id string) (*Project, error) {
	project, err := bolt.GetProjectByID(id)
	if err != nil {
		return nil, err
	}
	err = bolt.db.DeleteStruct(project)
	return project, err
}

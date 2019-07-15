package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/rs/xid"
)

type ProjectState string

const Active ProjectState = "active"
const Deleted ProjectState = "deleted"
const Archived ProjectState = "archived"
const Pending ProjectState = "pending"

// Project defines a simple project fields
type Project struct {
	ID        xid.ID `storm:"id"`
	Name      string
	Icon      string
	Age       int
	Seed      *Seed
	Author    string
	Version   string
	State     ProjectState
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newDefault(name string) (*Project, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().Unix())

	index := rand.Intn(len(icons))
	return &Project{
		ID:        xid.New(),
		Name:      name,
		Icon:      string(icons[index]),
		Author:    hostname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Age:       time.Now().Year(),
		State:     Pending,
		Version:   "0.0.1",
	}, nil
}

func (project *Project) plant(where string, seed *Seed) error {
	return seed.plant(where, project)
}

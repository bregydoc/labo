package main

import "time"

type Project struct {
	ProjectName string
	Author      string
	Version     string
	CreatedAt   time.Time
}

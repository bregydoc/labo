package main

import "time"

type Project struct {
	ProjectName string
	Age         int
	Author      string
	Version     string
	CreatedAt   time.Time
}

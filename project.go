package main

import "time"

// Project defines a simple project fields
type Project struct {
	ProjectName string
	Age         int
	Author      string
	Version     string
	CreatedAt   time.Time
}

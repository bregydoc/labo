package main

import (
	"time"

	"github.com/rs/xid"
)

// Project defines a simple project fields
type Project struct {
	ID        xid.ID
	Name      string
	Icon      string
	Age       int
	Template  Template
	Author    string
	Version   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

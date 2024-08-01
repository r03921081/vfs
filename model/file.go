package model

import "time"

type File struct {
	Name        string
	Description string
	Created     time.Time
}

func NewFile(name string, description string) *File {
	return &File{
		Name:        name,
		Description: description,
		Created:     time.Now().UTC(),
	}
}

package model

import "time"

type File struct {
	Name        string
	Description string
	CreatedAt   time.Time
}

func NewFile(name string, description string) *File {
	return &File{
		Name:        name,
		Description: description,
		CreatedAt:   time.Now().UTC(),
	}
}

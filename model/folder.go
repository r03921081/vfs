package model

import "time"

type Folder struct {
	Name        string
	Description string
	Files       map[string]*File
	CreatedAt   time.Time
}

func NewFolder(name string, description string) *Folder {
	return &Folder{
		Name:        name,
		Description: description,
		Files:       map[string]*File{},
		CreatedAt:   time.Now().UTC(),
	}
}

package model

import "time"

type Folder struct {
	Name        string
	Description string
	Files       map[string]*File
	Created     time.Time
}

func NewFolder(name string, description string) *Folder {
	return &Folder{
		Name:        name,
		Description: description,
		Files:       map[string]*File{},
		Created:     time.Now().UTC(),
	}
}

func (f *Folder) GetFiles() map[string]*File {
	return f.Files
}

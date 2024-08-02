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

func (f *Folder) GetName() string {
	return f.Name
}

func (f *Folder) GetCreated() time.Time {
	return f.Created
}

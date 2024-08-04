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

func (f *Folder) GetFiles() map[string]*File {
	return f.Files
}

func (f *Folder) SetFiles(files map[string]*File) {
	f.Files = files
}

func (f *Folder) GetFile(name string) *File {
	return f.Files[name]
}

func (f *Folder) SetFile(file *File) {
	f.Files[file.Name] = file
}

func (f *Folder) IsFileExist(name string) bool {
	_, ok := f.Files[name]
	return ok
}

func (f *Folder) DeleteFile(name string) {
	delete(f.Files, name)
}

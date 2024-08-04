package model

type FileSystem struct {
	Folders map[string]*Folder
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		Folders: map[string]*Folder{},
	}
}

func (fs *FileSystem) GetFolders() map[string]*Folder {
	return fs.Folders
}

func (fs *FileSystem) IsFolderExist(name string) bool {
	_, ok := fs.Folders[name]
	return ok
}

func (fs *FileSystem) GetFolder(name string) *Folder {
	return fs.Folders[name]
}

func (fs *FileSystem) SetFolder(folder *Folder) {
	fs.Folders[folder.Name] = folder
}

func (fs *FileSystem) DeleteFolder(name string) {
	delete(fs.Folders, name)
}

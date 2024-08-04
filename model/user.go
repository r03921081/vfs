package model

type User struct {
	Name       string
	FileSystem *FileSystem
}

func NewUser(name string) *User {
	return &User{
		Name:       name,
		FileSystem: NewFileSystem(),
	}
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetFileSystem() *FileSystem {
	return u.FileSystem
}

func (u *User) SetFileSystem(fileSystem *FileSystem) {
	u.FileSystem = fileSystem
}

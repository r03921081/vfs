package model

type User struct {
	Name    string
	Folders map[string]*Folder
}

func NewUser(name string) *User {
	return &User{
		Name:    name,
		Folders: map[string]*Folder{},
	}
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetFolders() map[string]*Folder {
	return u.Folders
}

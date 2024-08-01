package service

import "r03921081/vfs/repository"

func init() {
	UserService = NewUserService()
	FolderService = NewFolderService()
}

var (
	Register = repository.UserRepository.Register

	CreateFolder = repository.FolderRepository.Create
	DeleteFolder = repository.FolderRepository.Delete
	ListFolders  = repository.FolderRepository.List
	RenameFolder = repository.FolderRepository.Rename
)

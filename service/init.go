package service

import "r03921081/vfs/repository"

func init() {
	UserService = NewUserService()
	FolderService = NewFolderService()
	FileService = NewFileService()
}

var (
	Register = repository.UserRepository.Register

	CreateFolder = repository.FolderRepository.Create
	DeleteFolder = repository.FolderRepository.Delete
	ListFolders  = repository.FolderRepository.List
	RenameFolder = repository.FolderRepository.Rename

	CreateFile = repository.FileRepository.Create
	DeleteFile = repository.FileRepository.Delete
	ListFiles  = repository.FileRepository.List
)

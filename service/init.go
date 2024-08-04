package service

import "r03921081/vfs/repository"

func init() {
	UserService = NewUserService()
	FolderService = NewFolderService()
	FileService = NewFileService()
}

var (
	FlushUserCache = repository.FlushUserCache

	IsUserExist = repository.UserCacheRepository.IsUserExist
	SetUser     = repository.UserCacheRepository.SetUser

	IsUserFolderExist = repository.UserCacheRepository.IsUserFolderExist
	SetUserFolder     = repository.UserCacheRepository.SetUserFolder
	DeleteUserFolder  = repository.UserCacheRepository.DeleteUserFolder
	GetUserFolder     = repository.UserCacheRepository.GetUserFolder
	GetUserFolders    = repository.UserCacheRepository.GetUserFolders

	IsUserFileExist = repository.UserCacheRepository.IsUserFileExist
	SetUserFile     = repository.UserCacheRepository.SetUserFile
	GetUserFile     = repository.UserCacheRepository.GetUserFile
	GetUserFiles    = repository.UserCacheRepository.GetUserFiles
	DeleteUserFile  = repository.UserCacheRepository.DeleteUserFile
)

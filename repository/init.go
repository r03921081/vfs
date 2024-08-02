package repository

import "r03921081/vfs/common"

func init() {
	UserRepository = NewUserRepository()
	FolderRepository = NewFolderRepository()
	FileRepository = NewFileRepository()
}

var (
	CacheStorage = common.Cache
)

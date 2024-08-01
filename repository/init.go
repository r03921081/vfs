package repository

import "r03921081/vfs/common"

func init() {
	UserRepository = NewUserRepository()
	FolderRepository = NewFolderRepository()
}

var (
	CacheStorage = common.Cache
)

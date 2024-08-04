package repository

import (
	"fmt"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"

	"github.com/patrickmn/go-cache"
)

var userCache *cache.Cache = NewUserCache()

func NewUserCache() *cache.Cache {
	return cache.New(cache.NoExpiration, cache.NoExpiration)
}

func FlushUserCache() {
	userCache.Flush()
}

type IUserCacheRepository interface {
	SetUser(username string, user *model.User)
	GetUser(username string) (*model.User, common.ICodeError)
	IsUserExist(username string) bool
	GetUserFolders(username string) map[string]*model.Folder
	IsUserFolderExist(username string, folderName string) bool
	GetUserFolder(username string, folderName string) *model.Folder
	SetUserFolder(username string, folder *model.Folder)
	DeleteUserFolder(username string, folderName string)
	GetUserFiles(username string, folderName string) map[string]*model.File
	IsUserFileExist(username string, folderName string, fileName string) bool
	GetUserFile(username string, folderName string, fileName string) *model.File
	SetUserFile(username string, folderName string, file *model.File)
	DeleteUserFile(username string, folderName string, fileName string)
}

var UserCacheRepository IUserCacheRepository = NewUserCacheRepository()

type userCacheRepository struct{}

func NewUserCacheRepository() IUserCacheRepository {
	return &userCacheRepository{}
}

func (u *userCacheRepository) SetUser(username string, user *model.User) {
	userCache.Set(username, user, cache.NoExpiration)
}

func (u *userCacheRepository) GetUser(username string) (*model.User, common.ICodeError) {
	user, ok := userCache.Get(username)
	if !ok {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	return user.(*model.User), nil
}

func (u *userCacheRepository) IsUserExist(username string) bool {
	_, ok := userCache.Get(username)
	return ok
}

func (u *userCacheRepository) GetUserFolders(username string) map[string]*model.Folder {
	user, err := u.GetUser(username)
	if err != nil {
		return nil
	}
	return user.GetFileSystem().GetFolders()
}

func (u *userCacheRepository) IsUserFolderExist(username string, folderName string) bool {
	user, err := u.GetUser(username)
	if err != nil {
		return false
	}
	return user.GetFileSystem().IsFolderExist(folderName)
}

func (u *userCacheRepository) GetUserFolder(username string, folderName string) *model.Folder {
	user, err := u.GetUser(username)
	if err != nil {
		return nil
	}
	return user.GetFileSystem().GetFolder(folderName)
}

func (u *userCacheRepository) SetUserFolder(username string, folder *model.Folder) {
	user, err := u.GetUser(username)
	if err != nil {
		return
	}
	user.GetFileSystem().SetFolder(folder)
}

func (u *userCacheRepository) DeleteUserFolder(username string, folderName string) {
	user, err := u.GetUser(username)
	if err != nil {
		return
	}
	user.GetFileSystem().DeleteFolder(folderName)
}

func (u *userCacheRepository) GetUserFiles(username string, folderName string) map[string]*model.File {
	user, err := u.GetUser(username)
	if err != nil {
		return nil
	}
	return user.GetFileSystem().GetFolder(folderName).GetFiles()
}

func (u *userCacheRepository) IsUserFileExist(username string, folderName string, fileName string) bool {
	user, err := u.GetUser(username)
	if err != nil {
		return false
	}
	return user.GetFileSystem().GetFolder(folderName).IsFileExist(fileName)
}

func (u *userCacheRepository) GetUserFile(username string, folderName string, fileName string) *model.File {
	user, err := u.GetUser(username)
	if err != nil {
		return nil
	}
	return user.GetFileSystem().GetFolder(folderName).GetFile(fileName)
}

func (u *userCacheRepository) SetUserFile(username string, folderName string, file *model.File) {
	user, err := u.GetUser(username)
	if err != nil {
		return
	}
	user.GetFileSystem().GetFolder(folderName).SetFile(file)
}

func (u *userCacheRepository) DeleteUserFile(username string, folderName string, fileName string) {
	user, err := u.GetUser(username)
	if err != nil {
		return
	}
	user.GetFileSystem().GetFolder(folderName).DeleteFile(fileName)
}

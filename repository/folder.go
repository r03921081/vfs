package repository

import (
	"fmt"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"

	"github.com/patrickmn/go-cache"
)

type IFolderRepository interface {
	Create(username string, folder *model.Folder) (*model.Folder, common.ICodeError)
	Delete(username, folderName string) common.ICodeError
	List(username, sortby, orderby string) ([]*model.Folder, common.ICodeError)
	Rename(username, oldFolderName, newFolderName string) common.ICodeError
}

var FolderRepository IFolderRepository

type folderRepository struct{}

func NewFolderRepository() IFolderRepository {
	return &folderRepository{}
}

func (r *folderRepository) Create(username string, folder *model.Folder) (*model.Folder, common.ICodeError) {
	user, ok := CacheStorage.Get(username)
	if !ok {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	u := user.(*model.User)
	folders := u.GetFolders()
	if _, ok := folders[folder.Name]; ok {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, folder.Name))
	}
	folders[folder.Name] = folder
	u.Folders = folders
	CacheStorage.Set(username, u, cache.DefaultExpiration)
	return folder, nil
}

func (r *folderRepository) Delete(username string, folderName string) common.ICodeError {
	user, ok := CacheStorage.Get(username)
	if !ok {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	u := user.(*model.User)
	folders := u.GetFolders()
	if _, ok := folders[folderName]; !ok {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, folderName))
	}
	delete(folders, folderName)
	u.Folders = folders
	CacheStorage.Set(username, u, cache.DefaultExpiration)
	return nil
}

func (r *folderRepository) List(username, sortby, orderby string) ([]*model.Folder, common.ICodeError) {
	user, ok := CacheStorage.Get(username)
	if !ok {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	u := user.(*model.User)
	folders := []*model.Folder{}
	for _, v := range u.GetFolders() {
		folders = append(folders, v)
	}
	if len(folders) == 0 {
		return folders, nil
	}

	folders = sortFolders(folders, sortby, orderby)

	return folders, nil
}

func (r *folderRepository) Rename(username string, oldFolderName, newFolderName string) common.ICodeError {
	user, ok := CacheStorage.Get(username)
	if !ok {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	u := user.(*model.User)
	folders := u.GetFolders()
	if _, ok := folders[oldFolderName]; !ok {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, oldFolderName))
	}
	if _, ok := folders[newFolderName]; ok {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, newFolderName))
	}
	newFolder := model.NewFolder(newFolderName, folders[oldFolderName].Description)
	newFolder.Files = folders[oldFolderName].Files
	folders[newFolderName] = newFolder
	delete(folders, oldFolderName)
	u.Folders = folders
	CacheStorage.Set(username, u, cache.DefaultExpiration)
	return nil
}

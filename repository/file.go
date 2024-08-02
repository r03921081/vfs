package repository

import (
	"fmt"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"

	"github.com/patrickmn/go-cache"
)

type IFileRepository interface {
	Create(username, folderName string, file *model.File) (*model.File, common.ICodeError)
	Delete(username, folderName, fileName string) common.ICodeError
	List(username, folderName, sortby, orderby string) ([]*model.File, common.ICodeError)
}

var FileRepository IFileRepository

type fileRepository struct{}

func NewFileRepository() IFileRepository {
	return &fileRepository{}
}

func (r *fileRepository) Create(username, folderName string, file *model.File) (*model.File, common.ICodeError) {
	user, ok := CacheStorage.Get(username)
	if !ok {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	u := user.(*model.User)
	folders := u.GetFolders()
	if _, ok := folders[folderName]; !ok {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, folderName))
	}
	folder := folders[folderName]
	if _, ok := folder.Files[file.Name]; ok {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, file.Name))
	}
	folder.Files[file.Name] = file
	u.Folders[folderName] = folder
	CacheStorage.Set(username, u, cache.DefaultExpiration)
	return file, nil
}

func (r *fileRepository) Delete(username, folderName string, fileName string) common.ICodeError {
	user, ok := CacheStorage.Get(username)
	if !ok {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	u := user.(*model.User)
	folders := u.GetFolders()
	if _, ok := folders[folderName]; !ok {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, folderName))
	}
	folder := folders[folderName]
	if _, ok := folder.Files[fileName]; !ok {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, fileName))
	}
	delete(folder.Files, fileName)
	u.Folders[folderName] = folder
	CacheStorage.Set(username, u, cache.DefaultExpiration)
	return nil
}

func (r *fileRepository) List(username, folderName, sortby, orderby string) ([]*model.File, common.ICodeError) {
	user, ok := CacheStorage.Get(username)
	if !ok {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	u := user.(*model.User)
	folders := u.GetFolders()
	if _, ok := folders[folderName]; !ok {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, folderName))
	}
	folder := folders[folderName]
	files := []*model.File{}
	for _, file := range folder.Files {
		files = append(files, file)
	}
	if len(files) == 0 {
		return files, nil
	}

	files = sortItems(files, sortby, orderby)

	return files, nil
}

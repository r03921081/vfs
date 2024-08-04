package service

import (
	"fmt"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
)

type IFileService interface {
	Create(username, folderName, fileName, description string) (*model.File, common.ICodeError)
	Delete(username, folderName, fileName string) common.ICodeError
	List(username, folderName, sortby, orderby string) ([]*model.File, common.ICodeError)
}

var FileService IFileService

type fileService struct{}

func NewFileService() IFileService {
	return &fileService{}
}

func (s *fileService) Create(username, folderName, fileName, description string) (*model.File, common.ICodeError) {
	if !IsUserExist(username) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	if !IsUserFolderExist(username, folderName) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, folderName))
	}
	if IsUserFileExist(username, folderName, fileName) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, fileName))
	}

	file := model.NewFile(fileName, description)
	SetUserFile(username, folderName, file)
	return file, nil
}

func (s *fileService) Delete(username, folderName string, fileName string) common.ICodeError {
	if !IsUserExist(username) {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	if !IsUserFolderExist(username, folderName) {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, folderName))
	}
	if !IsUserFileExist(username, folderName, fileName) {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, fileName))
	}
	DeleteUserFile(username, folderName, fileName)
	return nil
}

func (s *fileService) List(username, folderName, sortby, orderby string) ([]*model.File, common.ICodeError) {
	if !IsUserExist(username) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	if !IsUserFolderExist(username, folderName) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, folderName))
	}
	files := []*model.File{}
	for _, file := range GetUserFiles(username, folderName) {
		files = append(files, file)
	}
	if len(files) == 0 {
		return files, nil
	}

	files = sortItems(files, sortby, orderby)

	return files, nil
}

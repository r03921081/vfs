package service

import (
	"fmt"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
)

type IFolderService interface {
	Create(username, folderName, description string) (*model.Folder, common.ICodeError)
	Delete(username, folderName string) common.ICodeError
	List(username, sortby, orderby string) ([]*model.Folder, common.ICodeError)
	Rename(username, oldFolderName, newFolderName string) common.ICodeError
}

var FolderService IFolderService

type folderService struct{}

func NewFolderService() IFolderService {
	return &folderService{}
}

func (s *folderService) Create(username, folderName, description string) (*model.Folder, common.ICodeError) {
	if !IsUserExist(username) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	if IsUserFolderExist(username, folderName) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, folderName))
	}

	folder := model.NewFolder(folderName, description)
	SetUserFolder(username, folder)
	return folder, nil
}

func (s *folderService) Delete(username, folderName string) common.ICodeError {
	if !IsUserExist(username) {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	if !IsUserFolderExist(username, folderName) {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, folderName))
	}
	DeleteUserFolder(username, folderName)
	return nil
}

func (s *folderService) List(username, sortby, orderby string) ([]*model.Folder, common.ICodeError) {
	if !IsUserExist(username) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	folders := []*model.Folder{}
	for _, folder := range GetUserFolders(username) {
		folders = append(folders, folder)
	}
	if len(folders) == 0 {
		return folders, nil
	}

	folders = sortItems(folders, sortby, orderby)

	return folders, nil
}

func (s *folderService) Rename(username, oldFolderName, newFolderName string) common.ICodeError {
	if !IsUserExist(username) {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}
	if !IsUserFolderExist(username, oldFolderName) {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, oldFolderName))
	}
	if IsUserFolderExist(username, newFolderName) {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, newFolderName))
	}
	oldFolder := GetUserFolder(username, oldFolderName)
	newFolder := model.NewFolder(newFolderName, oldFolder.Description)
	newFolder.SetFiles(oldFolder.GetFiles())

	SetUserFolder(username, newFolder)
	DeleteUserFolder(username, oldFolderName)
	return nil
}

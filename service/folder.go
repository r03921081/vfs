package service

import (
	"r03921081/vfs/common"
	"r03921081/vfs/model"
)

type IFolderService interface {
	Create(username string, folder *model.Folder) (*model.Folder, common.ICodeError)
	Delete(username, folderName string) common.ICodeError
	List(username, sortby, orderby string) ([]*model.Folder, common.ICodeError)
	Rename(username, oldFolderName, newFolderName string) common.ICodeError
}

var FolderService IFolderService

type folderService struct{}

func NewFolderService() IFolderService {
	return &folderService{}
}

func (s *folderService) Create(username string, folder *model.Folder) (*model.Folder, common.ICodeError) {
	return CreateFolder(username, folder)
}

func (s *folderService) Delete(username, folderName string) common.ICodeError {
	return DeleteFolder(username, folderName)
}

func (s *folderService) List(username, sortby, orderby string) ([]*model.Folder, common.ICodeError) {
	return ListFolders(username, sortby, orderby)
}

func (s *folderService) Rename(username, oldFolderName, newFolderName string) common.ICodeError {
	return RenameFolder(username, oldFolderName, newFolderName)
}

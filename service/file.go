package service

import (
	"r03921081/vfs/common"
	"r03921081/vfs/model"
)

type IFileService interface {
	Create(username, folderName string, file *model.File) (*model.File, common.ICodeError)
	Delete(username, folderName, fileName string) common.ICodeError
	List(username, folderName, sortby, orderby string) ([]*model.File, common.ICodeError)
}

var FileService IFileService

type fileService struct{}

func NewFileService() IFileService {
	return &fileService{}
}

func (s *fileService) Create(username, folderName string, file *model.File) (*model.File, common.ICodeError) {
	return CreateFile(username, folderName, file)
}

func (s *fileService) Delete(username, folderName string, fileName string) common.ICodeError {
	return DeleteFile(username, folderName, fileName)
}

func (s *fileService) List(username, folderName, sortby, orderby string) ([]*model.File, common.ICodeError) {
	return ListFiles(username, folderName, sortby, orderby)
}

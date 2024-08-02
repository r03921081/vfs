package service

import (
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fileService_Create(t *testing.T) {
	fileService := NewFileService()
	name := "user1"
	folder1 := "folder1"
	description1 := "description1"
	file1 := "file1"

	// Create file successfully
	CreateFile = func(username string, folderName string, file *model.File) (*model.File, common.ICodeError) {
		return file, nil
	}
	file, err := fileService.Create(name, folder1, model.NewFile(file1, description1))
	assert.Nil(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, file1, file.Name)
	assert.Equal(t, description1, file.Description)

	// Create file failed
	CreateFile = func(username string, folderName string, file *model.File) (*model.File, common.ICodeError) {
		return nil, common.NewCodeError(constant.ErrMsgHasAlreadyExisted)
	}
	file, err = fileService.Create(name, folder1, model.NewFile(file1, description1))
	assert.NotNil(t, err)
	assert.Equal(t, constant.ErrMsgHasAlreadyExisted, err.ErrorMessage())
	assert.Nil(t, file)
}

func Test_fileService_Delete(t *testing.T) {
	fileService := NewFileService()
	name := "user1"
	folder1 := "folder1"
	file1 := "file1"

	// Delete file successfully
	DeleteFile = func(username string, folderName string, fileName string) common.ICodeError {
		return nil
	}
	err := fileService.Delete(name, folder1, file1)
	assert.Nil(t, err)

	// Delete file failed
	DeleteFile = func(username string, folderName string, fileName string) common.ICodeError {
		return common.NewCodeError(constant.ErrMsgDoesNotExist)
	}
	err = fileService.Delete(name, folder1, file1)
	assert.NotNil(t, err)
	assert.Equal(t, constant.ErrMsgDoesNotExist, err.ErrorMessage())
}

func Test_fileService_List(t *testing.T) {
	fileService := NewFileService()
	name := "user1"
	folder1 := "folder1"
	file1 := "file1"

	// List files successfully
	ListFiles = func(username string, folderName string, sortby string, orderby string) ([]*model.File, common.ICodeError) {
		return []*model.File{model.NewFile(file1, "")}, nil
	}
	files, err := fileService.List(name, folder1, constant.SortCreated, constant.OrderAsc)
	assert.Nil(t, err)
	assert.NotNil(t, files)
	assert.Equal(t, 1, len(files))
	assert.Equal(t, file1, files[0].Name)

	// List files failed
	ListFiles = func(username string, folderName string, sortby string, orderby string) ([]*model.File, common.ICodeError) {
		return nil, common.NewCodeError(constant.ErrMsgDoesNotExist)
	}
	files, err = fileService.List(name, folder1, constant.SortCreated, constant.OrderAsc)
	assert.NotNil(t, err)
	assert.Equal(t, constant.ErrMsgDoesNotExist, err.ErrorMessage())
	assert.Nil(t, files)
}

package service

import (
	"fmt"
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

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return false
	}

	// User does not exist
	file, err := fileService.Create(name, folder1, file1, description1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, name), err.ErrorMessage())
	assert.Nil(t, file)

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return true
	}

	// Mock IsUserFolderExist
	IsUserFolderExist = func(username, folderName string) bool {
		return false
	}

	// Create file failed because folder does not exist
	file, err = fileService.Create(name, folder1, file1, description1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, folder1), err.ErrorMessage())
	assert.Nil(t, file)

	// Mock IsUserFolderExist
	IsUserFolderExist = func(username, folderName string) bool {
		return true
	}

	// Mock IsUserFileExist
	IsUserFileExist = func(username, folderName, fileName string) bool {
		return false
	}

	// Create file successfully
	file, err = fileService.Create(name, folder1, file1, description1)
	assert.Nil(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, file1, file.Name)

	// Mock IsUserFileExist
	IsUserFileExist = func(username, folderName, fileName string) bool {
		return true
	}

	// Create file failed because file already exists
	file, err = fileService.Create(name, folder1, file1, description1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, file1), err.ErrorMessage())
	assert.Nil(t, file)
}

func Test_fileService_Delete(t *testing.T) {
	fileService := NewFileService()
	name := "user1"
	folder1 := "folder1"
	file1 := "file1"

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return false
	}

	// User does not exist
	err := fileService.Delete(name, folder1, file1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, name), err.ErrorMessage())

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return true
	}

	// Mock IsUserFolderExist
	IsUserFolderExist = func(username, folderName string) bool {
		return false
	}

	// Folder does not exist
	err = fileService.Delete(name, folder1, file1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, folder1), err.ErrorMessage())

	// Mock IsUserFolderExist
	IsUserFolderExist = func(username, folderName string) bool {
		return true
	}

	// Mock IsUserFileExist
	IsUserFileExist = func(username, folderName, fileName string) bool {
		return false
	}

	// File does not exist
	err = fileService.Delete(name, folder1, file1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, file1), err.ErrorMessage())

	// Mock IsUserFileExist
	IsUserFileExist = func(username, folderName, fileName string) bool {
		return true
	}

	// Delete file successfully
	err = fileService.Delete(name, folder1, file1)
	assert.Nil(t, err)
}

func Test_fileService_List(t *testing.T) {
	fileService := NewFileService()
	name := "user1"
	folder1 := "folder1"
	file1 := "file1"
	sortby := constant.SortName
	orderby := constant.OrderAsc

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return false
	}

	// User does not exist
	files, err := fileService.List(name, folder1, sortby, orderby)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, name), err.ErrorMessage())
	assert.Nil(t, files)

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return true
	}

	// Mock IsUserFolderExist
	IsUserFolderExist = func(username, folderName string) bool {
		return false
	}

	// Folder does not exist
	files, err = fileService.List(name, folder1, sortby, orderby)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, folder1), err.ErrorMessage())
	assert.Nil(t, files)

	// Mock IsUserFolderExist
	IsUserFolderExist = func(username, folderName string) bool {
		return true
	}

	// Mock GetUserFiles
	GetUserFiles = func(username, folderName string) map[string]*model.File {
		return nil
	}

	// List files successfully
	files, err = fileService.List(name, folder1, sortby, orderby)
	assert.Nil(t, err)
	assert.NotNil(t, files)
	assert.Equal(t, 0, len(files))

	// Mock GetUserFiles
	GetUserFiles = func(username, folderName string) map[string]*model.File {
		return map[string]*model.File{
			file1: {
				Name: file1,
			},
		}
	}

	// List files successfully
	files, err = fileService.List(name, folder1, sortby, orderby)
	assert.Nil(t, err)
	assert.NotNil(t, files)
	assert.Equal(t, 1, len(files))
}

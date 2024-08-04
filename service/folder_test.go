package service

import (
	"fmt"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_folderService_Create(t *testing.T) {
	folderService := NewFolderService()
	username := "user1"
	folder1 := "folder1"
	description1 := "description1"

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return false
	}

	// User does not exist
	folder, err := folderService.Create(username, folder1, description1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, username), err.ErrorMessage())
	assert.Nil(t, folder)

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return true
	}

	// Mock IsUserFolderExist
	IsUserFolderExist = func(username, folderName string) bool {
		return false
	}

	// Create folder successfully
	folder, err = folderService.Create(username, folder1, description1)
	assert.Nil(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder1, folder.Name)
	assert.Equal(t, description1, folder.Description)

	// Mock IsUserFolderExist
	IsUserFolderExist = func(username, folderName string) bool {
		return true
	}

	// Duplicate folder
	folder, err = folderService.Create(username, folder1, description1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, folder1), err.ErrorMessage())
	assert.Nil(t, folder)
}

func Test_folderService_Delete(t *testing.T) {
	folderService := NewFolderService()
	username := "user1"
	folder1 := "folder1"

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return false
	}

	// User does not exist
	err := folderService.Delete(username, folder1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, username), err.ErrorMessage())

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return true
	}

	// Mock IsUserFolderExist
	IsUserFolderExist = func(username, folderName string) bool {
		return false
	}

	// Folder does not exist
	err = folderService.Delete(username, folder1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, folder1), err.ErrorMessage())

	// Mock IsUserFolderExist
	IsUserFolderExist = func(username, folderName string) bool {
		return true
	}

	// Delete folder successfully
	err = folderService.Delete(username, folder1)
	assert.Nil(t, err)
}

func Test_folderService_List(t *testing.T) {
	folderService := NewFolderService()
	username := "user1"
	folder1 := "folder1"
	folder2 := "folder2"
	sortby := constant.SortName
	orderby := constant.OrderAsc

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return false
	}

	// User does not exist
	folders, err := folderService.List(username, sortby, orderby)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, username), err.ErrorMessage())
	assert.Nil(t, folders)

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return true
	}

	// Mock GetUserFolders
	GetUserFolders = func(username string) map[string]*model.Folder {
		return map[string]*model.Folder{
			folder1: {
				Name: folder1,
			},
			folder2: {
				Name: folder2,
			},
		}
	}

	// List folders successfully
	folders, err = folderService.List(username, sortby, orderby)
	assert.Nil(t, err)
	assert.NotNil(t, folders)
	assert.Equal(t, 2, len(folders))
	assert.Equal(t, folder1, folders[0].Name)
	assert.Equal(t, folder2, folders[1].Name)
}

func Test_folderService_Rename(t *testing.T) {
	folderService := NewFolderService()
	username := "user1"
	oldFolder := "oldFolder"
	newFolder := "newFolder"
	description := "description"

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return false
	}

	// User does not exist
	err := folderService.Rename(username, oldFolder, newFolder)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, username), err.ErrorMessage())

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return true
	}

	// Mock IsUserFolderExist, oldFolder not exist, newFolder not exist
	IsUserFolderExist = func(username, folderName string) bool {
		return false
	}

	// Rename folder failed because oldFolder does not exist
	err = folderService.Rename(username, oldFolder, newFolder)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, oldFolder), err.ErrorMessage())

	// Mock IsUserFolderExist, oldFolder exist, newFolder exist
	IsUserFolderExist = func(username, folderName string) bool {
		return true
	}

	// Rename folder failed because newFolder has already existed
	err = folderService.Rename(username, oldFolder, newFolder)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, newFolder), err.ErrorMessage())

	// Mock IsUserFolderExist, oldFolder exist, newFolder not exist
	IsUserFolderExist = func(username, folderName string) bool {
		return folderName == oldFolder
	}

	// Mock GetUserFolder
	GetUserFolder = func(username, folderName string) *model.Folder {
		return model.NewFolder(folderName, description)
	}

	// Rename folder successfully
	err = folderService.Rename(username, oldFolder, newFolder)
	assert.Nil(t, err)
}

package service

import (
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_folderService_Create(t *testing.T) {
	folderService := NewFolderService()
	name := "user1"
	folder1 := "folder1"
	description1 := "description1"

	// Create folder successfully
	CreateFolder = func(username string, folder *model.Folder) (*model.Folder, common.ICodeError) {
		return folder, nil
	}
	folder, err := folderService.Create(name, model.NewFolder(folder1, description1))
	assert.Nil(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder1, folder.Name)
	assert.Equal(t, description1, folder.Description)

	// Create folder failed
	CreateFolder = func(username string, folder *model.Folder) (*model.Folder, common.ICodeError) {
		return nil, common.NewCodeError(constant.ErrMsgHasAlreadyExisted)
	}
	folder, err = folderService.Create(name, model.NewFolder(folder1, description1))
	assert.NotNil(t, err)
	assert.Equal(t, constant.ErrMsgHasAlreadyExisted, err.ErrorMessage())
	assert.Nil(t, folder)
}

func Test_folderService_Delete(t *testing.T) {
	folderService := NewFolderService()
	name := "user1"
	folder1 := "folder1"

	// Delete folder successfully
	DeleteFolder = func(username string, folderName string) common.ICodeError {
		return nil
	}
	err := folderService.Delete(name, folder1)
	assert.Nil(t, err)

	// Delete folder failed
	DeleteFolder = func(username string, folderName string) common.ICodeError {
		return common.NewCodeError(constant.ErrMsgDoesNotExist)
	}
	err = folderService.Delete(name, folder1)
	assert.NotNil(t, err)
	assert.Equal(t, constant.ErrMsgDoesNotExist, err.ErrorMessage())
}

func Test_folderService_List(t *testing.T) {
	folderService := NewFolderService()
	name := "user1"
	folder1 := "folder1"

	// List folders successfully
	ListFolders = func(username string, sortby string, orderby string) ([]*model.Folder, common.ICodeError) {
		return []*model.Folder{model.NewFolder(folder1, "")}, nil
	}
	folders, err := folderService.List(name, constant.SortCreated, constant.OrderAsc)
	assert.Nil(t, err)
	assert.NotNil(t, folders)
	assert.Equal(t, 1, len(folders))
	assert.Equal(t, folder1, folders[0].Name)

	// List folders failed
	ListFolders = func(username string, sortby string, orderby string) ([]*model.Folder, common.ICodeError) {
		return nil, common.NewCodeError(constant.ErrMsgDoesNotExist)
	}
	folders, err = folderService.List(name, constant.SortCreated, constant.OrderAsc)
	assert.NotNil(t, err)
	assert.Equal(t, constant.ErrMsgDoesNotExist, err.ErrorMessage())
	assert.Nil(t, folders)
}

func Test_folderService_Rename(t *testing.T) {
	folderService := NewFolderService()
	name := "user1"
	folder1 := "folder1"
	newFolder := "newFolder"

	// Rename folder successfully
	RenameFolder = func(username string, oldFolderName string, newFolderName string) common.ICodeError {
		return nil
	}
	err := folderService.Rename(name, folder1, newFolder)
	assert.Nil(t, err)

	// Rename folder failed
	RenameFolder = func(username string, oldFolderName string, newFolderName string) common.ICodeError {
		return common.NewCodeError(constant.ErrMsgDoesNotExist)
	}
	err = folderService.Rename(name, folder1, newFolder)
	assert.NotNil(t, err)
	assert.Equal(t, constant.ErrMsgDoesNotExist, err.ErrorMessage())
}

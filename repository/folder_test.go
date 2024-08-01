package repository

import (
	"fmt"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_folderRepository_Create(t *testing.T) {
	CacheStorage = common.NewCache()
	userRepository := NewUserRepository()
	folderRepository := NewFolderRepository()

	username1 := "user1"

	// User does not exist
	folder1 := model.NewFolder("folder1", "description1")
	folder, err := folderRepository.Create(username1, folder1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, username1), err.ErrorMessage())
	assert.Nil(t, folder)

	// Registered user1 successfully
	user1 := model.NewUser(username1)
	err = userRepository.Register(user1)
	assert.Nil(t, err)

	folder, err = folderRepository.Create(username1, folder1)
	assert.Nil(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder1.Name, folder.Name)
	assert.Equal(t, folder1.Description, folder.Description)
	assert.Equal(t, folder1.Files, folder.Files)
}

func Test_folderRepository_Delete(t *testing.T) {
	CacheStorage = common.NewCache()
	userRepository := NewUserRepository()
	folderRepository := NewFolderRepository()

	username1 := "user1"
	folderName1 := "folder1"

	// User does not exist
	err := folderRepository.Delete(username1, folderName1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, username1), err.ErrorMessage())

	// Registered user1 successfully
	user1 := model.NewUser(username1)
	err = userRepository.Register(user1)
	assert.Nil(t, err)

	// Folder does not exist
	err = folderRepository.Delete(username1, folderName1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, folderName1), err.ErrorMessage())

	// Registered user1 and folder1 successfully
	folder1 := model.NewFolder(folderName1, "description1")
	folder, err := folderRepository.Create(username1, folder1)
	assert.Nil(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder1.Name, folder.Name)
	assert.Equal(t, folder1.Description, folder.Description)
	assert.Equal(t, folder1.Files, folder.Files)

	// Delete folder1 successfully
	err = folderRepository.Delete(username1, folderName1)
	assert.Nil(t, err)
}

func Test_folderRepository_List(t *testing.T) {
	CacheStorage = common.NewCache()
	userRepository := NewUserRepository()
	folderRepository := NewFolderRepository()

	username1 := "user1"
	username2 := "user2"
	folderName1 := "folder1"
	folderName2 := "folder2"

	// User1 does not exist
	folders, err := folderRepository.List(username1, constant.SortCreated, constant.OrderDesc)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, username1), err.ErrorMessage())
	assert.Nil(t, folders)

	// Registered user1 successfully
	user1 := model.NewUser(username1)
	err = userRepository.Register(user1)
	assert.Nil(t, err)

	// Folder does not exist
	folders, err = folderRepository.List(username1, constant.SortCreated, constant.OrderDesc)
	assert.Nil(t, err)
	assert.Len(t, folders, 0)

	// Registered user1 and folder1 successfully
	folder1 := model.NewFolder(folderName1, "description1")
	folder, err := folderRepository.Create(username1, folder1)
	assert.Nil(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder1.Name, folder.Name)
	assert.Equal(t, folder1.Description, folder.Description)
	assert.Equal(t, folder1.Files, folder.Files)

	// List folder1 successfully
	folders, err = folderRepository.List(username1, constant.SortCreated, constant.OrderDesc)
	assert.Nil(t, err)
	assert.Len(t, folders, 1)
	assert.Equal(t, folder1.Name, folders[0].Name)
	assert.Equal(t, folder1.Description, folders[0].Description)
	assert.Equal(t, folder1.Files, folders[0].Files)

	// Registered user1 and folder2 successfully
	folder2 := model.NewFolder(folderName2, "description2")
	folder, err = folderRepository.Create(username1, folder2)
	assert.Nil(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder2.Name, folder.Name)
	assert.Equal(t, folder2.Description, folder.Description)
	assert.Equal(t, folder2.Files, folder.Files)

	// List folder1 and folder2 successfully with sort name and order asc
	folders, err = folderRepository.List(username1, constant.SortName, constant.OrderAsc)
	assert.Nil(t, err)
	assert.Len(t, folders, 2)
	assert.Equal(t, folder1.Name, folders[0].Name)
	assert.Equal(t, folder1.Description, folders[0].Description)
	assert.Equal(t, folder1.Files, folders[0].Files)
	assert.Equal(t, folder2.Name, folders[1].Name)
	assert.Equal(t, folder2.Description, folders[1].Description)
	assert.Equal(t, folder2.Files, folders[1].Files)

	// List folder1 and folder2 successfully with sort name and order desc
	folders, err = folderRepository.List(username1, constant.SortName, constant.OrderDesc)
	assert.Nil(t, err)
	assert.Len(t, folders, 2)
	assert.Equal(t, folder2.Name, folders[0].Name)
	assert.Equal(t, folder2.Description, folders[0].Description)
	assert.Equal(t, folder2.Files, folders[0].Files)
	assert.Equal(t, folder1.Name, folders[1].Name)
	assert.Equal(t, folder1.Description, folders[1].Description)
	assert.Equal(t, folder1.Files, folders[1].Files)

	// List folder1 and folder2 successfully with sort created and order asc
	folders, err = folderRepository.List(username1, constant.SortCreated, constant.OrderAsc)
	assert.Nil(t, err)
	assert.Len(t, folders, 2)
	assert.Equal(t, folder1.Name, folders[0].Name)
	assert.Equal(t, folder1.Description, folders[0].Description)
	assert.Equal(t, folder1.Files, folders[0].Files)
	assert.Equal(t, folder2.Name, folders[1].Name)
	assert.Equal(t, folder2.Description, folders[1].Description)
	assert.Equal(t, folder2.Files, folders[1].Files)

	// List folder1 and folder2 successfully with sort created and order desc
	folders, err = folderRepository.List(username1, constant.SortCreated, constant.OrderDesc)
	assert.Nil(t, err)
	assert.Len(t, folders, 2)
	assert.Equal(t, folder2.Name, folders[0].Name)
	assert.Equal(t, folder2.Description, folders[0].Description)
	assert.Equal(t, folder2.Files, folders[0].Files)
	assert.Equal(t, folder1.Name, folders[1].Name)
	assert.Equal(t, folder1.Description, folders[1].Description)
	assert.Equal(t, folder1.Files, folders[1].Files)

	// User2 does not exist
	folders, err = folderRepository.List(username2, constant.SortCreated, constant.OrderDesc)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, username2), err.ErrorMessage())
	assert.Nil(t, folders)

	// Registered user2 successfully
	user2 := model.NewUser(username2)
	err = userRepository.Register(user2)
	assert.Nil(t, err)

	// Folder does not exist
	folders, err = folderRepository.List(username2, constant.SortCreated, constant.OrderDesc)
	assert.Nil(t, err)
	assert.Len(t, folders, 0)
}

func Test_folderRepository_Rename(t *testing.T) {
	CacheStorage = common.NewCache()
	userRepository := NewUserRepository()
	folderRepository := NewFolderRepository()

	username1 := "user1"
	oldFolderName1 := "oldFolderName1"
	newFolderName1 := "newfoldername1"

	// User does not exist
	err := folderRepository.Rename(username1, oldFolderName1, newFolderName1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, username1), err.ErrorMessage())

	// Registered user1 successfully
	user1 := model.NewUser(username1)
	err = userRepository.Register(user1)
	assert.Nil(t, err)

	// Folder does not exist
	err = folderRepository.Rename(username1, oldFolderName1, newFolderName1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, oldFolderName1), err.ErrorMessage())

	// Registered user1 and oldFolderName1 successfully
	folder1 := model.NewFolder(oldFolderName1, "description1")
	folder, err := folderRepository.Create(username1, folder1)
	assert.Nil(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder1.Name, folder.Name)
	assert.Equal(t, folder1.Description, folder.Description)
	assert.Equal(t, folder1.Files, folder.Files)

	// Rename oldFolderName1 to newFolderName1 successfully
	err = folderRepository.Rename(username1, oldFolderName1, newFolderName1)
	assert.Nil(t, err)

	folders, err := folderRepository.List(username1, constant.SortCreated, constant.OrderDesc)
	assert.Nil(t, err)
	assert.Len(t, folders, 1)
	assert.Equal(t, newFolderName1, folders[0].Name)
	assert.Equal(t, folder1.Description, folders[0].Description)
	assert.Equal(t, folder1.Files, folders[0].Files)

	// Registered user1 and oldFolderName1 successfully
	folder, err = folderRepository.Create(username1, folder1)
	assert.Nil(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder1.Name, folder.Name)
	assert.Equal(t, folder1.Description, folder.Description)
	assert.Equal(t, folder1.Files, folder.Files)

	// Rename newFolderName1 to oldFolderName1 failed
	err = folderRepository.Rename(username1, newFolderName1, oldFolderName1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, oldFolderName1), err.ErrorMessage())
}

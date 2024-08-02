package repository

import (
	"fmt"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fileRepository_Create(t *testing.T) {
	CacheStorage = common.NewCache()
	userRepository := NewUserRepository()
	folderRepository := NewFolderRepository()
	fileRepository := NewFileRepository()

	username1 := "user1"
	folder1 := model.NewFolder("folder1", "description1")

	// User does not exist
	file1 := model.NewFile("file1", "description1")
	file1, err := fileRepository.Create(username1, folder1.GetName(), file1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, username1), err.ErrorMessage())
	assert.Nil(t, file1)

	// Registered user1 successfully
	user1 := model.NewUser(username1)
	err = userRepository.Register(user1)
	assert.Nil(t, err)

	// Folder does not exist
	file1, err = fileRepository.Create(username1, folder1.GetName(), file1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, folder1.GetName()), err.ErrorMessage())
	assert.Nil(t, file1)

	// Registered user1 and folder1 successfully
	folder, err := folderRepository.Create(username1, folder1)
	assert.Nil(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder1.Name, folder.Name)
	assert.Equal(t, folder1.Description, folder.Description)
	assert.Equal(t, folder1.Files, folder.Files)

	// Registered user1 and folder1 and file1 successfully
	file1 = model.NewFile("file1", "description1")
	file1, err = fileRepository.Create(username1, folder1.GetName(), file1)
	assert.Nil(t, err)
	assert.NotNil(t, file1)
	assert.Equal(t, file1.Name, file1.Name)
	assert.Equal(t, file1.Description, file1.Description)

	// Duplicate file
	file1 = model.NewFile("file1", "description1")
	file1, err = fileRepository.Create(username1, folder1.GetName(), file1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, "file1"), err.ErrorMessage())
	assert.Nil(t, file1)
}

func Test_fileRepository_Delete(t *testing.T) {
	CacheStorage = common.NewCache()
	userRepository := NewUserRepository()
	folderRepository := NewFolderRepository()
	fileRepository := NewFileRepository()

	username1 := "user1"
	folder1 := model.NewFolder("folder1", "description1")

	// User does not exist
	file1 := model.NewFile("file1", "description1")
	err := fileRepository.Delete(username1, folder1.GetName(), file1.GetName())
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, username1), err.ErrorMessage())

	// Registered user1 successfully
	user1 := model.NewUser(username1)
	err = userRepository.Register(user1)
	assert.Nil(t, err)

	// Folder does not exist
	err = fileRepository.Delete(username1, folder1.GetName(), file1.GetName())
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, folder1.GetName()), err.ErrorMessage())

	// Registered user1 and folder1 successfully
	folder, err := folderRepository.Create(username1, folder1)
	assert.Nil(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder1.Name, folder.Name)
	assert.Equal(t, folder1.Description, folder.Description)
	assert.Equal(t, folder1.Files, folder.Files)

	// File does not exist
	err = fileRepository.Delete(username1, folder1.GetName(), file1.GetName())
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, file1.GetName()), err.ErrorMessage())

	// Registered user1 and folder1 and file1 successfully
	file1 = model.NewFile("file1", "description1")
	file1, err = fileRepository.Create(username1, folder1.GetName(), file1)
	assert.Nil(t, err)
	assert.NotNil(t, file1)
	assert.Equal(t, file1.Name, file1.Name)
	assert.Equal(t, file1.Description, file1.Description)

	// Delete file1 successfully
	err = fileRepository.Delete(username1, folder1.GetName(), file1.GetName())
	assert.Nil(t, err)
	assert.Equal(t, 0, len(folder1.Files))
}

func Test_fileRepository_List(t *testing.T) {
	CacheStorage = common.NewCache()
	userRepository := NewUserRepository()
	folderRepository := NewFolderRepository()
	fileRepository := NewFileRepository()

	username1 := "user1"
	folder1 := model.NewFolder("folder1", "description1")
	file1 := model.NewFile("file1", "description1")
	file2 := model.NewFile("file2", "description2")

	// User does not exist
	files, err := fileRepository.List(username1, folder1.GetName(), constant.SortName, constant.OrderAsc)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, username1), err.ErrorMessage())
	assert.Nil(t, files)

	// Registered user1 successfully
	user1 := model.NewUser(username1)
	err = userRepository.Register(user1)
	assert.Nil(t, err)

	// Folder does not exist
	files, err = fileRepository.List(username1, folder1.GetName(), constant.SortName, constant.OrderAsc)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgDoesNotExist, folder1.GetName()), err.ErrorMessage())
	assert.Nil(t, files)

	// Registered user1 and folder1 successfully
	folder, err := folderRepository.Create(username1, folder1)
	assert.Nil(t, err)
	assert.NotNil(t, folder)
	assert.Equal(t, folder1.Name, folder.Name)
	assert.Equal(t, folder1.Description, folder.Description)
	assert.Equal(t, folder1.Files, folder.Files)

	// File does not exist
	files, err = fileRepository.List(username1, folder1.GetName(), constant.SortName, constant.OrderAsc)
	assert.Nil(t, err)
	assert.Len(t, files, 0)

	// Registered user1 and folder1 and file1 successfully
	file1, err = fileRepository.Create(username1, folder1.GetName(), file1)
	assert.Nil(t, err)
	assert.NotNil(t, file1)
	assert.Equal(t, file1.Name, file1.Name)
	assert.Equal(t, file1.Description, file1.Description)

	// List file1 successfully
	files, err = fileRepository.List(username1, folder1.GetName(), constant.SortName, constant.OrderAsc)
	assert.Nil(t, err)
	assert.Len(t, files, 1)
	assert.Equal(t, file1.Name, files[0].Name)
	assert.Equal(t, file1.Description, files[0].Description)

	// Registered user1 and folder1 and file2 successfully
	file2, err = fileRepository.Create(username1, folder1.GetName(), file2)
	assert.Nil(t, err)
	assert.NotNil(t, file2)
	assert.Equal(t, file2.Name, file2.Name)
	assert.Equal(t, file2.Description, file2.Description)

	// List file1 and file2 successfully
	files, err = fileRepository.List(username1, folder1.GetName(), constant.SortName, constant.OrderAsc)
	assert.Nil(t, err)
	assert.Len(t, files, 2)
	assert.Equal(t, file1.Name, files[0].Name)
	assert.Equal(t, file1.Description, files[0].Description)
	assert.Equal(t, file2.Name, files[1].Name)
	assert.Equal(t, file2.Description, files[1].Description)
}

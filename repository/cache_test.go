package repository

import (
	"fmt"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAndGetUser(t *testing.T) {
	FlushUserCache()
	userCacheRepository := NewUserCacheRepository()

	username := "testuser"
	user := model.NewUser(username)

	// User does not exist
	cachedUser, err := userCacheRepository.GetUser(username)
	assert.NotNil(t, err)
	assert.Equal(t, err, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username)))
	assert.Nil(t, cachedUser)

	// User exists
	userCacheRepository.SetUser(username, user)
	cachedUser, err = userCacheRepository.GetUser(username)
	assert.Nil(t, err)
	assert.Equal(t, user, cachedUser)
}

func TestIsUserExist(t *testing.T) {
	FlushUserCache()
	userCacheRepository := NewUserCacheRepository()

	username := "testuser"
	user := model.NewUser(username)

	userCacheRepository.SetUser(username, user)

	exists := userCacheRepository.IsUserExist(username)
	assert.True(t, exists)

	notExists := userCacheRepository.IsUserExist("nonexistentuser")
	assert.False(t, notExists)
}

func TestGetUserFolders(t *testing.T) {
	FlushUserCache()
	userCacheRepository := NewUserCacheRepository()

	username := "testuser"
	user := model.NewUser(username)

	userCacheRepository.SetUser(username, user)
	folders := userCacheRepository.GetUserFolders(username)

	assert.NotNil(t, folders)
	assert.Equal(t, 0, len(folders))
}

func TestSetAndGetUserFolder(t *testing.T) {
	FlushUserCache()
	userCacheRepository := NewUserCacheRepository()

	username := "testuser"
	user := model.NewUser(username)

	userCacheRepository.SetUser(username, user)

	folder := &model.Folder{Name: "folder1"}
	userCacheRepository.SetUserFolder(username, folder)

	retrievedFolder := userCacheRepository.GetUserFolder(username, "folder1")
	assert.Equal(t, folder, retrievedFolder)
}

func TestIsUserFolderExist(t *testing.T) {
	FlushUserCache()
	userCacheRepository := NewUserCacheRepository()

	username := "testuser"
	user := model.NewUser(username)

	userCacheRepository.SetUser(username, user)

	folder := &model.Folder{Name: "folder1"}
	userCacheRepository.SetUserFolder(username, folder)

	exists := userCacheRepository.IsUserFolderExist(username, "folder1")
	assert.True(t, exists)

	notExists := userCacheRepository.IsUserFolderExist(username, "nonexistentfolder")
	assert.False(t, notExists)
}

func TestSetAndGetUserFile(t *testing.T) {
	FlushUserCache()
	userCacheRepository := NewUserCacheRepository()

	username := "testuser"
	user := model.NewUser(username)

	userCacheRepository.SetUser(username, user)

	folder := &model.Folder{Name: "folder1", Files: make(map[string]*model.File)}
	userCacheRepository.SetUserFolder(username, folder)

	file := &model.File{Name: "file1"}
	userCacheRepository.SetUserFile(username, "folder1", file)

	retrievedFile := userCacheRepository.GetUserFile(username, "folder1", "file1")
	assert.Equal(t, file, retrievedFile)

	files := userCacheRepository.GetUserFiles(username, "folder1")
	assert.Equal(t, 1, len(files))
	assert.Equal(t, file, files["file1"])
}

func TestIsUserFileExist(t *testing.T) {
	FlushUserCache()
	userCacheRepository := NewUserCacheRepository()

	username := "testuser"
	user := model.NewUser(username)

	userCacheRepository.SetUser(username, user)

	folder := &model.Folder{Name: "folder1", Files: make(map[string]*model.File)}
	userCacheRepository.SetUserFolder(username, folder)

	file := &model.File{Name: "file1"}
	userCacheRepository.SetUserFile(username, "folder1", file)

	exists := userCacheRepository.IsUserFileExist(username, "folder1", "file1")
	assert.True(t, exists)

	notExists := userCacheRepository.IsUserFileExist(username, "folder1", "nonexistentfile")
	assert.False(t, notExists)
}

func TestDeleteUserFolder(t *testing.T) {
	FlushUserCache()
	userCacheRepository := NewUserCacheRepository()

	username := "testuser"
	user := model.NewUser(username)

	userCacheRepository.SetUser(username, user)

	folder := &model.Folder{Name: "folder1"}
	userCacheRepository.SetUserFolder(username, folder)

	userCacheRepository.DeleteUserFolder(username, "folder1")

	exists := userCacheRepository.IsUserFolderExist(username, "folder1")
	assert.False(t, exists)
}

func TestDeleteUserFile(t *testing.T) {
	FlushUserCache()
	userCacheRepository := NewUserCacheRepository()

	username := "testuser"
	user := model.NewUser(username)

	userCacheRepository.SetUser(username, user)

	folder := &model.Folder{Name: "folder1", Files: make(map[string]*model.File)}
	userCacheRepository.SetUserFolder(username, folder)

	file := &model.File{Name: "file1"}
	userCacheRepository.SetUserFile(username, "folder1", file)

	userCacheRepository.DeleteUserFile(username, "folder1", "file1")

	exists := userCacheRepository.IsUserFileExist(username, "folder1", "file1")
	assert.False(t, exists)
}

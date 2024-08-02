package controller

import (
	"bytes"
	"fmt"
	"os"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"r03921081/vfs/util"
	"testing"

	"gotest.tools/assert"
)

func Test_folderController_Create(t *testing.T) {
	folderController := NewFolderController()
	name := "user1"
	folderName1 := "folder1"
	description1 := "description1"

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	CreateFolder = func(username string, folder *model.Folder) (*model.Folder, common.ICodeError) {
		return folder, nil
	}

	folderController.Create(name, folderName1, description1)

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.MsgCreateSuccessfully, folderName1)
	assert.Equal(t, expectedOutput+"\n", buf.String())
}

func Test_folderController_Create_with_error(t *testing.T) {
	folderController := NewFolderController()
	name := "user1"
	folderName1 := "folder1"
	description1 := "description1"

	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	CreateFolder = func(username string, folder *model.Folder) (*model.Folder, common.ICodeError) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}

	folderController.Create(name, folderName1, description1)

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.ErrMsgDoesNotExist, name)
	assert.Equal(t, constant.PrefixError+expectedOutput+"\n", buf.String())
}

func Test_folderController_Delete(t *testing.T) {
	folderController := NewFolderController()
	name := "user1"
	folderName1 := "folder1"

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	DeleteFolder = func(username string, folderName string) common.ICodeError {
		return nil
	}

	folderController.Delete(name, folderName1)

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.MsgDeleteSuccessfully, folderName1)
	assert.Equal(t, expectedOutput+"\n", buf.String())
}

func Test_folderController_Delete_with_error(t *testing.T) {
	folderController := NewFolderController()
	name := "user1"
	folderName1 := "folder1"

	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	DeleteFolder = func(username string, folderName string) common.ICodeError {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}

	folderController.Delete(name, folderName1)

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.ErrMsgDoesNotExist, name)
	assert.Equal(t, constant.PrefixError+expectedOutput+"\n", buf.String())
}

func Test_folderController_List(t *testing.T) {
	folderController := NewFolderController()
	name := "user1"
	folderName1 := "folder1"
	folder := model.NewFolder(folderName1, "")

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ListFolders = func(username string, sortby string, orderby string) ([]*model.Folder, common.ICodeError) {
		return []*model.Folder{folder}, nil
	}

	folderController.List(name, constant.SortCreated, constant.OrderAsc)

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := util.FormatFolders([]*model.Folder{folder}, name)
	assert.Equal(t, expectedOutput+"\n", buf.String())
}

func Test_folderController_List_with_error(t *testing.T) {
	folderController := NewFolderController()
	name := "user1"

	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	ListFolders = func(username string, sortby string, orderby string) ([]*model.Folder, common.ICodeError) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}

	folderController.List(name, constant.SortCreated, constant.OrderAsc)

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.ErrMsgDoesNotExist, name)
	assert.Equal(t, constant.PrefixError+expectedOutput+"\n", buf.String())
}

func Test_folderController_List_with_warning(t *testing.T) {
	folderController := NewFolderController()
	name := "user1"

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ListFolders = func(username string, sortby string, orderby string) ([]*model.Folder, common.ICodeError) {
		return []*model.Folder{}, nil
	}

	folderController.List(name, constant.SortCreated, constant.OrderAsc)

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.WarningMsgDoesNotHaveAnyFolders, name)
	assert.Equal(t, constant.PrefixWaring+expectedOutput+"\n", buf.String())
}

func Test_folderController_Rename(t *testing.T) {
	folderController := NewFolderController()
	name := "user1"
	folderName1 := "folder1"
	folderName2 := "folder2"

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	RenameFolder = func(username string, oldFolderName string, newFolderName string) common.ICodeError {
		return nil
	}

	folderController.Rename(name, folderName1, folderName2)

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.MsgRenameSuccessfully, folderName1, folderName2)
	assert.Equal(t, expectedOutput+"\n", buf.String())
}

func Test_folderController_Rename_with_error(t *testing.T) {
	folderController := NewFolderController()
	name := "user1"
	folderName1 := "folder1"
	folderName2 := "folder2"

	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	RenameFolder = func(username string, oldFolderName string, newFolderName string) common.ICodeError {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}

	folderController.Rename(name, folderName1, folderName2)

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.ErrMsgDoesNotExist, name)
	assert.Equal(t, constant.PrefixError+expectedOutput+"\n", buf.String())
}

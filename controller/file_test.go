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
	"time"

	"gotest.tools/assert"
)

func Test_fileController_Create(t *testing.T) {
	fileController := NewFileController()
	name := "user1"
	folderName1 := "folder1"
	description1 := "description1"
	file1 := "file1"

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	CreateFile = func(username string, folderName string, fileName string, description string) (*model.File, common.ICodeError) {
		return model.NewFile(file1, description), nil
	}

	fileController.Create(name, folderName1, file1, description1)

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.MsgCreateFileSuccessfully, file1, name, folderName1)
	assert.Equal(t, expectedOutput+"\n", buf.String())
}

func Test_fileController_Create_with_wrong_format(t *testing.T) {
	fileController := NewFileController()
	name := "user1"
	folderName1 := "folder1"
	description1 := "description1"
	file1 := "# file1"

	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	fileController.Create(name, folderName1, file1, description1)

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.ErrMsgContainInvalidChars, file1)
	assert.Equal(t, constant.PrefixError+expectedOutput+"\n", buf.String())
}

func Test_fileController_Create_with_error(t *testing.T) {
	fileController := NewFileController()
	name := "user1"
	folderName1 := "folder1"
	description1 := "description1"
	file1 := "file1"

	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	CreateFile = func(username string, folderName string, fileName string, description string) (*model.File, common.ICodeError) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}

	fileController.Create(name, folderName1, file1, description1)

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.ErrMsgDoesNotExist, name)
	assert.Equal(t, constant.PrefixError+expectedOutput+"\n", buf.String())
}

func Test_fileController_Delete(t *testing.T) {
	fileController := NewFileController()
	name := "user1"
	folderName1 := "folder1"
	file1 := "file1"

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	DeleteFile = func(username string, folderName string, fileName string) common.ICodeError {
		return nil
	}

	fileController.Delete(name, folderName1, file1)

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.MsgDeleteSuccessfully, file1)
	assert.Equal(t, expectedOutput+"\n", buf.String())
}

func Test_fileController_Delete_with_error(t *testing.T) {
	fileController := NewFileController()
	name := "user1"
	folderName1 := "folder1"
	file1 := "file1"

	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	DeleteFile = func(username string, folderName string, fileName string) common.ICodeError {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}

	fileController.Delete(name, folderName1, file1)

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.ErrMsgDoesNotExist, name)
	assert.Equal(t, constant.PrefixError+expectedOutput+"\n", buf.String())
}

func Test_fileController_List(t *testing.T) {
	fileController := NewFileController()
	name := "user1"
	folderName1 := "folder1"
	fileName1 := "file1"
	file1 := &model.File{
		Name:        fileName1,
		Description: "description1",
		Created:     time.Date(2024, time.August, 1, 15, 0, 10, 0, time.UTC),
	}

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ListFiles = func(username string, folderName string, sortby string, orderby string) ([]*model.File, common.ICodeError) {
		return []*model.File{file1}, nil
	}

	fileController.List(name, folderName1, constant.SortCreated, constant.OrderAsc)

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := util.FormatFiles([]*model.File{file1}, name, folderName1)
	assert.Equal(t, expectedOutput+"\n", buf.String())
}

func Test_fileController_List_with_error(t *testing.T) {
	fileController := NewFileController()
	name := "user1"
	folderName1 := "folder1"

	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	ListFiles = func(username string, folderName string, sortby string, orderby string) ([]*model.File, common.ICodeError) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgDoesNotExist, username))
	}

	fileController.List(name, folderName1, constant.SortCreated, constant.OrderAsc)
	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.ErrMsgDoesNotExist, name)
	assert.Equal(t, constant.PrefixError+expectedOutput+"\n", buf.String())
}

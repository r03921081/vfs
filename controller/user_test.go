package controller

import (
	"bytes"
	"fmt"
	"os"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"testing"

	"gotest.tools/assert"
)

func Test_userController_Register(t *testing.T) {
	userController := NewUserController()
	name := "user1"

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Register = func(username string) (*model.User, common.ICodeError) {
		return model.NewUser(name), nil
	}

	userController.Register(name)

	w.Close()
	os.Stdout = originalStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.MsgAddSuccessfully, name)
	assert.Equal(t, expectedOutput+"\n", buf.String())
}

func Test_userController_Register_with_wrong_format(t *testing.T) {
	userController := NewUserController()
	name := "#user1"

	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	userController.Register(name)

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.ErrMsgContainInvalidChars, name)
	assert.Equal(t, constant.PrefixError+expectedOutput+"\n", buf.String())
}

func Test_userController_Register_with_error(t *testing.T) {
	userController := NewUserController()
	name := "user1"

	originalStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	Register = func(username string) (*model.User, common.ICodeError) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, username))
	}

	userController.Register(name)

	w.Close()
	os.Stderr = originalStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expectedOutput := fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, name)
	assert.Equal(t, constant.PrefixError+expectedOutput+"\n", buf.String())
}

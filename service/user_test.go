package service

import (
	"fmt"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_userService_Register(t *testing.T) {
	userService := NewUserService()
	name := "user1"

	// Registered user successfully
	Register = func(user *model.User) common.ICodeError {
		return nil
	}
	expectedUser := model.NewUser(name)

	user, err := userService.Register(name)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)

	// Registered user successfully
	Register = func(user *model.User) common.ICodeError {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, user.Name))
	}

	user, err = userService.Register(name)
	assert.Nil(t, user)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, name), err.ErrorMessage())
}

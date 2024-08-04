package service

import (
	"fmt"
	"r03921081/vfs/constant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_userService_Register(t *testing.T) {
	FlushUserCache()

	userService := NewUserService()
	name := "user1"

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return false
	}

	// Register successfully
	user, err := userService.Register(name)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, name, user.GetName())

	// Mock IsUserExist
	IsUserExist = func(username string) bool {
		return true
	}

	// Register failed
	user, err = userService.Register(name)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, name), err.ErrorMessage())
	assert.Nil(t, user)
}

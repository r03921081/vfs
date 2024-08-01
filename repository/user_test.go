package repository

import (
	"fmt"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_userRepository_Register(t *testing.T) {
	CacheStorage = common.NewCache()
	userRepository := NewUserRepository()

	// Registered user1 successfully
	user1 := model.NewUser("user1")
	err := userRepository.Register(user1)
	assert.Nil(t, err)

	// Repeated registration of user1 failed
	err = userRepository.Register(user1)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, user1.GetName()), err.ErrorMessage())

	// Registered user2 successfully
	user2 := model.NewUser("user2")
	err = userRepository.Register(user2)
	assert.Nil(t, err)
}

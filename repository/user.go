package repository

import (
	"fmt"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"

	"github.com/patrickmn/go-cache"
)

type IUserRepository interface {
	Register(user *model.User) common.ICodeError
}

var UserRepository IUserRepository

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (r *userRepository) Register(user *model.User) common.ICodeError {
	username := user.GetName()

	if _, ok := common.Cache.Get(username); ok {
		return common.NewCodeError(fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, username))
	}

	common.Cache.Set(username, user, cache.DefaultExpiration)
	return nil
}

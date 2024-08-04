package service

import (
	"fmt"
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
)

type IUserService interface {
	Register(username string) (*model.User, common.ICodeError)
}

var UserService IUserService

type userService struct{}

func NewUserService() *userService {
	return &userService{}
}

func (u *userService) Register(username string) (*model.User, common.ICodeError) {
	if IsUserExist(username) {
		return nil, common.NewCodeError(fmt.Sprintf(constant.ErrMsgHasAlreadyExisted, username))
	}

	user := model.NewUser(username)
	SetUser(username, user)
	return user, nil
}

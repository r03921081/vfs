package service

import (
	"r03921081/vfs/common"
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

func (u *userService) Register(name string) (*model.User, common.ICodeError) {
	user := model.NewUser(name)
	err := Register(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

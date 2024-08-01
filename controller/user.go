package controller

import (
	"fmt"
	"r03921081/vfs/constant"
)

type IUserController interface {
	Register(username string)
}

var UserController IUserController

type userController struct{}

func NewUserController() IUserController {
	return &userController{}
}

func (u *userController) Register(name string) {
	user, err := Register(name)
	if err != nil {
		PrintError(err.ErrorMessage())
		return
	}
	PrintSuccess(fmt.Sprintf(constant.MsgAddSuccessfully, user.GetName()))
}

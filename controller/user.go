package controller

import (
	"fmt"
	"r03921081/vfs/constant"
	"r03921081/vfs/util"
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
	if !util.IsValidInput(name, util.ValidName) {
		PrintError(fmt.Sprintf(constant.ErrMsgContainInvalidChars, name))
		return
	}
	user, err := Register(name)
	if err != nil {
		PrintError(err.ErrorMessage())
		return
	}
	PrintSuccess(fmt.Sprintf(constant.MsgAddSuccessfully, user.GetName()))
}

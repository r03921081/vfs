package controller

import (
	"r03921081/vfs/common"
	"r03921081/vfs/constant"
	"strings"
)

type ICommandController interface {
	Handle(command string)
}

var CommandController ICommandController

type commandController struct{}

func NewCommandController() ICommandController {
	return &commandController{}
}

func (c *commandController) Handle(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		common.Printer.PrintError(constant.ErrMsgUnrecognizedCommand)
		return
	}

	switch constant.Command(parts[0]) {
	case constant.CommandRegister:
		UserController.Register(parts[1])
	default:
		common.Printer.PrintError(constant.ErrMsgUnrecognizedCommand)
	}
}

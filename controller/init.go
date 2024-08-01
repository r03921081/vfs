package controller

import (
	"r03921081/vfs/common"
	"r03921081/vfs/service"
)

func init() {
	CommandController = NewCommandController()
	UserController = NewUserController()
}

var (
	PrintSuccess = common.Printer.PrintSuccess
	PrintWarning = common.Printer.PrintWarning
	PrintError   = common.Printer.PrintError
)

var (
	Register = service.UserService.Register
)

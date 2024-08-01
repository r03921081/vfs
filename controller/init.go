package controller

import (
	"r03921081/vfs/common"
	"r03921081/vfs/service"
)

func init() {
	CommandController = NewCommandController()
	UserController = NewUserController()
	FolderController = NewFolderController()
}

var (
	PrintSuccess = common.Printer.PrintSuccess
	PrintWarning = common.Printer.PrintWarning
	PrintError   = common.Printer.PrintError
)

var (
	Register = service.UserService.Register
)

var (
	CreateFolder = service.FolderService.Create
	DeleteFolder = service.FolderService.Delete
	ListFolders  = service.FolderService.List
	RenameFolder = service.FolderService.Rename
)

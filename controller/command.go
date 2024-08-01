package controller

import (
	"fmt"
	"r03921081/vfs/constant"
	"r03921081/vfs/util"
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
		PrintError(constant.ErrMsgUnrecognizedCommand)
		return
	}
	fmt.Println(parts)

	switch constant.Command(parts[0]) {
	case constant.CommandRegister:
		if len(parts) != 2 {
			PrintError(constant.PromptRegister)
			return
		}
		username := parts[1]
		UserController.Register(username)
	case constant.CommandCreateFolder:
		if len(parts) < 3 {
			PrintError(constant.PromptCreateFolder)
			return
		}
		FolderController.Create(parts[1], parts[2], strings.Join(parts[3:], " "))
	case constant.CommandDeleteFolder:
		if len(parts) != 3 {
			PrintError(constant.PromptDeleteFolder)
			return
		}
		username := parts[1]
		foldername := parts[2]
		FolderController.Delete(username, foldername)
	case constant.CommandListFolders:
		if len(parts) < 2 || len(parts) > 4 {
			PrintError(constant.PromptListFolders)
			return
		}
		sortby := constant.SortName
		orderby := constant.OrderAsc
		if len(parts) > 2 {
			sortby = parts[2]
		}
		if len(parts) > 3 {
			orderby = parts[3]
		}
		if !util.IsValidListParams(sortby, orderby) {
			PrintError(constant.PromptListFolders)
			return
		}
		username := parts[1]
		FolderController.List(username, sortby, orderby)
	case constant.CommandRenameFolder:
		if len(parts) != 4 {
			PrintError(constant.PromptRenameFolder)
			return
		}
		username := parts[1]
		oldfoldername := parts[2]
		newfoldername := parts[3]
		FolderController.Rename(username, oldfoldername, newfoldername)
	default:
		PrintError(constant.ErrMsgUnrecognizedCommand)
	}
}

package controller

import (
	"fmt"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"r03921081/vfs/util"
)

type IFolderController interface {
	Create(username string, foldername, description string)
	Delete(username, foldername string)
	List(username, sortby, orderby string)
	Rename(username, oldfoldername, foldername string)
}

var FolderController IFolderController

type folderController struct{}

func NewFolderController() IFolderController {
	return &folderController{}
}

func (c *folderController) Create(username string, foldername, description string) {
	folder := model.NewFolder(foldername, description)
	folder, err := CreateFolder(username, folder)
	if err != nil {
		PrintError(err.ErrorMessage())
		return
	}
	PrintSuccess(fmt.Sprintf(constant.MsgCreateSuccessfully, folder.Name))
}

func (c *folderController) Delete(username, foldername string) {
	err := DeleteFolder(username, foldername)
	if err != nil {
		PrintError(err.ErrorMessage())
		return
	}
	PrintSuccess(fmt.Sprintf(constant.MsgDeleteSuccessfully, foldername))
}

func (c *folderController) List(username, sortby, orderby string) {
	folders, err := ListFolders(username, sortby, orderby)
	if err != nil {
		PrintError(err.ErrorMessage())
		return
	}
	if len(folders) == 0 {
		PrintWarning(fmt.Sprintf(constant.WarningMsgDoesNotHaveAnyFolders, username))
		return
	}
	PrintSuccess(util.FormatFolders(folders, username))
}

func (c *folderController) Rename(username, oldfoldername, foldername string) {
	err := RenameFolder(username, oldfoldername, foldername)
	if err != nil {
		PrintError(err.ErrorMessage())
		return
	}
	PrintSuccess(fmt.Sprintf(constant.MsgRenameSuccessfully, oldfoldername, foldername))
}

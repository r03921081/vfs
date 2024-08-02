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

func (c *folderController) Create(username string, folderName, description string) {
	if !util.IsValidInput(folderName, util.ValidName) {
		PrintError(fmt.Sprintf(constant.ErrMsgContainInvalidChars, folderName))
		return
	}
	folder := model.NewFolder(folderName, description)
	folder, err := CreateFolder(username, folder)
	if err != nil {
		PrintError(err.ErrorMessage())
		return
	}
	PrintSuccess(fmt.Sprintf(constant.MsgCreateSuccessfully, folder.Name))
}

func (c *folderController) Delete(username, folderName string) {
	err := DeleteFolder(username, folderName)
	if err != nil {
		PrintError(err.ErrorMessage())
		return
	}
	PrintSuccess(fmt.Sprintf(constant.MsgDeleteSuccessfully, folderName))
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

func (c *folderController) Rename(username, oldFolderName, newFolderName string) {
	if !util.IsValidInput(newFolderName, util.ValidName) {
		PrintError(fmt.Sprintf(constant.ErrMsgContainInvalidChars, newFolderName))
		return
	}
	err := RenameFolder(username, oldFolderName, newFolderName)
	if err != nil {
		PrintError(err.ErrorMessage())
		return
	}
	PrintSuccess(fmt.Sprintf(constant.MsgRenameSuccessfully, oldFolderName, newFolderName))
}

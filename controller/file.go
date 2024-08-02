package controller

import (
	"fmt"
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"r03921081/vfs/util"
)

type IFileController interface {
	Create(username, folderName, fileName, description string)
	Delete(username, folderName, fileName string)
	List(username, folderName, sortby, orderby string)
}

var FileController IFileController

type fileController struct{}

func NewFileController() IFileController {
	return &fileController{}
}

func (c *fileController) Create(username, folderName, fileName, description string) {
	if !util.IsValidInput(fileName, util.ValidName) {
		PrintError(fmt.Sprintf(constant.ErrMsgContainInvalidChars, fileName))
		return
	}
	file := model.NewFile(fileName, description)
	file, err := CreateFile(username, folderName, file)
	if err != nil {
		PrintError(err.ErrorMessage())
		return
	}
	PrintSuccess(fmt.Sprintf(constant.MsgCreateFileSuccessfully, file.Name, username, folderName))
}

func (c *fileController) Delete(username, folderName, fileName string) {
	err := DeleteFile(username, folderName, fileName)
	if err != nil {
		PrintError(err.ErrorMessage())
		return
	}
	PrintSuccess(fmt.Sprintf(constant.MsgDeleteSuccessfully, fileName))
}

func (c *fileController) List(username, folderName, sortby, orderby string) {
	files, err := ListFiles(username, folderName, sortby, orderby)
	if err != nil {
		PrintError(err.ErrorMessage())
		return
	}
	if len(files) == 0 {
		PrintWarning(constant.WarningMsgTheFolderIsEmpty)
		return
	}
	PrintSuccess(util.FormatFiles(files, username, folderName))
}

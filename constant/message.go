package constant

import (
	"fmt"
)

var (
	PromptRegister     = "Usage: register [username]"
	PromptCreateFolder = "Usage: create-folder [username] [foldername] [description]?"
	PromptDeleteFolder = "Usage: delete-folder [username] [foldername]"
	PromptListFolders  = "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]"
	PromptRenameFolder = "Usage: rename-folder [username] [foldername] [new-folder-name]"
	PromptCreateFile   = "Usage: create-file [username] [foldername] [filename] [description]?"
	PromptDeleteFile   = "Usage: delete-file [username] [foldername] [filename]"
	PromptListFiles    = "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]"
)

var (
	MsgAddSuccessfully        = "Add %s successfully."
	MsgCreateSuccessfully     = "Create %s successfully."
	MsgDeleteSuccessfully     = "Delete %s successfully."
	MsgRenameSuccessfully     = "Rename %s to %s successfully."
	MsgCreateFileSuccessfully = "Create %s in %s/%s successfully."
)

var (
	WarningMsgDoesNotHaveAnyFolders = "The %s doesn't have any folders."
	WarningMsgTheFolderIsEmpty      = "The folder is empty."
)

var (
	PrefixError  = "Error: "
	PrefixWaring = "Warning: "

	ErrMsgUnrecognizedCommand = "Unrecognized command"

	ErrMsgHasAlreadyExisted   = "The %s has already existed."
	ErrMsgContainInvalidChars = "The %s contain invalid chars."
	ErrMsgDoesNotExist        = "The %s doesn't exist."

	ErrMsgCommandShouldNotBeLongerThan = fmt.Sprintf("The command should not be longer than %d chars.", MaxLengthCommand)
)

package constant

var (
	PromptRegister     = "Usage: register [username]"
	PromptCreateFolder = "Usage: create-folder [username] [foldername] [description]?"
	PromptDeleteFolder = "Usage: delete-folder [username] [foldername]"
	PromptListFolders  = "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]"
	PromptRenameFolder = "Usage: rename-folder [username] [foldername] [new-folder-name]"
)

var (
	MsgAddSuccessfully    = "Add %s successfully."
	MsgCreateSuccessfully = "Create %s successfully."
	MsgDeleteSuccessfully = "Delete %s successfully."
	MsgRenameSuccessfully = "Rename %s to %s successfully."
)

var (
	WarningMsgDoesNotHaveAnyFolders = "The %s doesn't have any folders."
)

var (
	PrefixError  = "Error: "
	PrefixWaring = "Warning: "

	ErrMsgUnrecognizedCommand = "Unrecognized command"

	ErrMsgHasAlreadyExisted   = "The %s has already existed."
	ErrMsgContainInvalidChars = "The %s contain invalid chars."
	ErrMsgDoesNotExist        = "The %s doesn't exist."
)

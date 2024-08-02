package constant

type Command string

const (
	CommandRegister Command = "register"

	CommandCreateFolder Command = "create-folder"
	CommandDeleteFolder Command = "delete-folder"
	CommandListFolders  Command = "list-folders"
	CommandRenameFolder Command = "rename-folder"

	CommandCreateFile Command = "create-file"
	CommandDeleteFile Command = "delete-file"
	CommandListFiles  Command = "list-files"
)

func (c Command) String() string {
	return string(c)
}

const (
	MaxLengthCommand     = 500
	MaxLengthName        = 16
	MaxLengthDescription = 100
)

package controller

import (
	"bytes"
	"os"
	"regexp"
	"testing"

	"gotest.tools/assert"
)

func Test_commandController_Handle(t *testing.T) {
	commandController := NewCommandController()

	stdoutR, stdoutW, _ := os.Pipe()
	stderrR, stderrW, _ := os.Pipe()

	originalStdout := os.Stdout
	originalStderr := os.Stderr

	os.Stdout = stdoutW
	os.Stderr = stderrW

	defer func() {
		os.Stdout = originalStdout
		os.Stderr = originalStderr
	}()

	commandController.Handle("register user1")
	commandController.Handle("register user2")
	commandController.Handle("create-folder user1 folder1")
	commandController.Handle("create-folder user2 folder1")
	commandController.Handle("create-folder user1 folder1")
	commandController.Handle("create-folder user1 folder2 this-is-folder-2")
	commandController.Handle("list-folders user1 --sort-name asc")
	commandController.Handle("list-folders user2")
	commandController.Handle("create-file user1 folder1 file1 this-is-file1")
	commandController.Handle("create-file user1 folder1 config a-config-file")
	commandController.Handle("create-file user1 folder1 config a-config-file")
	commandController.Handle("create-file user-abc folder-abc config a-config-file")
	commandController.Handle("list data")
	commandController.Handle("list-files user1 folder1 --sort a")
	commandController.Handle("list-files user1 folder1 --sort-name desc")

	stdoutW.Close()
	stderrW.Close()

	var stdoutBuf, stderrBuf bytes.Buffer
	_, _ = stdoutBuf.ReadFrom(stdoutR)
	_, _ = stderrBuf.ReadFrom(stderrR)

	timeRegexp := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`)

	// Standard output
	expectedStdout := `Add user1 successfully.
Add user2 successfully.
Create folder1 successfully.
Create folder1 successfully.
Create folder2 successfully.
folder1		<time>	user1
folder2	this-is-folder-2	<time>	user1
folder1		<time>	user2
Create file1 in user1/folder1 successfully.
Create config in user1/folder1 successfully.
file1	this-is-file1	<time>	folder1	user1
config	a-config-file	<time>	folder1	user1`

	output := timeRegexp.ReplaceAllString(stdoutBuf.String(), "<time>")
	assert.Equal(t, expectedStdout+"\n", output)

	// Standard error
	expectedStderr := `Error: The folder1 has already existed.
Error: The config has already existed.
Error: The user-abc doesn't exist.
Error: Unrecognized command
Error: Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]`

	assert.Equal(t, expectedStderr+"\n", stderrBuf.String())
}

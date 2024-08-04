# Virtual File System Side Project

This is a virtual file system with user and file management.

## Technologies Used

- **Go** v1.21
- **testify** v1.9.0
- **go-cache** v2.1.0+incompatible

## How to Run

```sh
go build
./vfs
```

## How to Use
### Register a user
```sh
Usage: register [username]
```
### Create a folder
```sh
Usage: create-folder [username] [foldername] [description]?
```
### Delete a folder
```sh
Usage: delete-folder [username] [foldername]
```
### List all folders
```sh
Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]
```
### Rename a folder
```sh
Usage: rename-folder [username] [foldername] [new-folder-name]
```
### Create a file
```sh
Usage: create-file [username] [foldername] [filename] [description]?
```
### Delete a file
```sh
Usage: delete-file [username] [foldername] [filename]
```
### List all files
```sh
Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]
```
### Exit the application
```sh
Usage: exit
```

## Introduction
1. The entry point is the command controller, responsible for recognizing and processing input commands.
2. The maximum lengths are defined as follows: MaxLengthCommand = 500, MaxLengthName = 16, and MaxLengthDescription = 100.
3. Commands can contain only the following characters: `a-z`, `A-Z`, `0-9`, and the symbols `+`, `-`, `*`, `/`, `_`, `@`, `[`, `]`, `(`, `)`, `{`, `}`, `.`, and spaces.
4. Names can only contain `a-z`, `A-Z`, `0-9` and the symbols `-`, `_` and `.`.
5. The file system data structure is as follows:
```go
type User struct {
	Name       string
	FileSystem *FileSystem
}
type FileSystem struct {
	Folders map[string]*Folder
}
type Folder struct {
	Name        string
	Description string
	Files       map[string]*File
	Created     time.Time
}
type File struct {
	Name        string
	Description string
	Created     time.Time
}
```
6. The system uses `userCache`, implemented via the `go-cache` library, to temporarily store data in memory. 
Here’s an example of the data structure:
```
{
  "Name": "JohnDoe",
  "FileSystem": {
    "Folders": {
      "folder1": {
        "Name": "Documents",
        "Description": "Personal Documents",
        "Files": {
          "file1": {
            "Name": "Resume.pdf",
            "Description": "My latest resume",
            "Created": "2023-08-01T10:00:00Z"
          },
          "file2": {
            "Name": "CoverLetter.pdf",
            "Description": "Cover letter for job application",
            "Created": "2023-08-01T11:00:00Z"
          }
        },
        "Created": "2023-08-01T09:00:00Z"
      },
      "folder2": {
        "Name": "Photos",
        "Description": "Family vacation photos",
        "Files": {
          "file1": {
            "Name": "Beach.jpg",
            "Description": "Photo of the beach",
            "Created": "2023-07-20T15:00:00Z"
          },
          "file2": {
            "Name": "Mountain.jpg",
            "Description": "Photo of the mountain",
            "Created": "2023-07-21T16:00:00Z"
          }
        },
        "Created": "2023-07-20T14:00:00Z"
      }
    }
  }
}
```
7. Interfaces are utilized to facilitate easy changes and accommodate different implementations.
8. All errors in the system are thrown as ICodeError and are uniformly printed at the controller layer.
9. The controller, service, and repository layers all include unit tests for their respective functions.
10. For testing the `command` controller, I used the examples provided in the assignment to perform tests.

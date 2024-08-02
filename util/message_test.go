package util

import (
	"r03921081/vfs/model"
	"testing"
	"time"
)

func TestFormatFolders(t *testing.T) {
	type args struct {
		folders  []*model.Folder
		username string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				folders: []*model.Folder{
					{
						Name:        "folder1",
						Description: "description1",
						Created:     time.Date(2024, time.August, 1, 15, 0, 10, 0, time.UTC),
					},
					{
						Name:        "folder2",
						Description: "description2",
						Created:     time.Date(2024, time.August, 2, 0, 0, 10, 0, time.UTC),
					},
				},
				username: "user1",
			},
			want: "folder1\tdescription1\t2024-08-01 15:00:10\tuser1\nfolder2\tdescription2\t2024-08-02 00:00:10\tuser1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatFolders(tt.args.folders, tt.args.username); got != tt.want {
				t.Errorf("FormatFolders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatFiles(t *testing.T) {
	type args struct {
		files      []*model.File
		username   string
		foldername string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				files: []*model.File{
					{
						Name:        "file1",
						Description: "description1",
						Created:     time.Date(2024, time.August, 1, 15, 0, 10, 0, time.UTC),
					},
					{
						Name:        "file2",
						Description: "description2",
						Created:     time.Date(2024, time.August, 2, 0, 0, 10, 0, time.UTC),
					},
				},
				username:   "user1",
				foldername: "folder1",
			},
			want: "file1\tdescription1\t2024-08-01 15:00:10\tfolder1\tuser1\nfile2\tdescription2\t2024-08-02 00:00:10\tfolder1\tuser1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatFiles(tt.args.files, tt.args.username, tt.args.foldername); got != tt.want {
				t.Errorf("FormatFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidListParams(t *testing.T) {
	type args struct {
		sortby  string
		orderby string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				sortby:  "name",
				orderby: "asc",
			},
			want: false,
		},
		{
			name: "case2",
			args: args{
				sortby:  "created",
				orderby: "ascdesc",
			},
			want: false,
		},
		{
			name: "case3",
			args: args{
				sortby:  "--sort-name",
				orderby: "desc",
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				sortby:  "--sort-created",
				orderby: "desc",
			},
			want: true,
		},
		{
			name: "case5",
			args: args{
				sortby:  "",
				orderby: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidListParams(tt.args.sortby, tt.args.orderby); got != tt.want {
				t.Errorf("IsValidListParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

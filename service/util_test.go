package service

import (
	"r03921081/vfs/model"
	"reflect"
	"testing"
	"time"
)

func Test_sortItems_folders(t *testing.T) {
	folders := []*model.Folder{
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
	}

	type args struct {
		folders []*model.Folder
		sortby  string
		orderby string
	}
	tests := []struct {
		name string
		args args
		want []*model.Folder
	}{
		{
			name: "case1",
			args: args{
				folders: folders,
				sortby:  "name",
				orderby: "asc",
			},
			want: folders,
		},
		{
			name: "case2",
			args: args{
				folders: folders,
				sortby:  "name",
				orderby: "desc",
			},
			want: []*model.Folder{folders[1], folders[0]},
		},
		{
			name: "case3",
			args: args{
				folders: folders,
				sortby:  "created",
				orderby: "asc",
			},
			want: folders,
		},
		{
			name: "case4",
			args: args{
				folders: folders,
				sortby:  "created",
				orderby: "desc",
			},
			want: []*model.Folder{folders[1], folders[0]},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortItems(tt.args.folders, tt.args.sortby, tt.args.orderby); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortFolders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortFiles(t *testing.T) {
	files := []*model.File{
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
	}

	type args struct {
		files   []*model.File
		sortby  string
		orderby string
	}
	tests := []struct {
		name string
		args args
		want []*model.File
	}{
		{
			name: "case1",
			args: args{
				files:   files,
				sortby:  "name",
				orderby: "asc",
			},
			want: files,
		},
		{
			name: "case2",
			args: args{
				files:   files,
				sortby:  "name",
				orderby: "desc",
			},
			want: []*model.File{files[1], files[0]},
		},
		{
			name: "case3",
			args: args{
				files:   files,
				sortby:  "created",
				orderby: "asc",
			},
			want: files,
		},
		{
			name: "case4",
			args: args{
				files:   files,
				sortby:  "created",
				orderby: "desc",
			},
			want: []*model.File{files[1], files[0]},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortItems(tt.args.files, tt.args.sortby, tt.args.orderby); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

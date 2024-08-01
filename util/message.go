package util

import (
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
)

func FormatFolders(folders []*model.Folder, username string) string {
	s := ""
	for i, v := range folders {
		s += v.Name + "\t" + v.Description + "\t" + v.Created.Format("2006-01-02 15:04:05") + "\t" + username
		if i < len(folders)-1 {
			s += "\n"
		}
	}
	return s
}

func FormatFiles(files []*model.File, username, foldername string) string {
	s := ""
	for i, v := range files {
		s += v.Name + "\t" + v.Description + "\t" + v.Created.Format("2006-01-02 15:04:05") + "\t" + foldername + "\t" + username
		if i < len(files)-1 {
			s += "\n"
		}
	}
	return s
}

func IsValidListParams(sortby, orderby string) bool {
	if sortby != constant.SortName && sortby != constant.SortCreated {
		return false
	}
	if orderby != constant.OrderAsc && orderby != constant.OrderDesc {
		return false
	}
	return true
}

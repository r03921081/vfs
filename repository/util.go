package repository

import (
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"sort"
)

func sortFolders(folders []*model.Folder, sortby, orderby string) []*model.Folder {
	if sortby != constant.SortCreated {
		sortby = constant.SortName
	}
	if orderby != constant.OrderDesc {
		orderby = constant.OrderAsc
	}

	if sortby == constant.SortCreated {
		if orderby == constant.OrderAsc {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].Created.Before(folders[j].Created)
			})
		} else {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].Created.After(folders[j].Created)
			})
		}
	} else {
		if orderby == constant.OrderAsc {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].Name < folders[j].Name
			})
		} else {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].Name > folders[j].Name
			})
		}
	}
	return folders
}

package service

import (
	"r03921081/vfs/constant"
	"r03921081/vfs/model"
	"sort"
)

func sortItems[T model.Sortable](items []T, sortby, orderby string) []T {
	if sortby != constant.SortName && sortby != constant.SortCreated {
		sortby = constant.SortName
	}
	if orderby != constant.OrderAsc && orderby != constant.OrderDesc {
		orderby = constant.OrderAsc
	}

	if sortby == constant.SortCreated {
		if orderby == constant.OrderAsc {
			sort.Slice(items, func(i, j int) bool {
				return items[i].GetCreated().Before(items[j].GetCreated())
			})
		} else {
			sort.Slice(items, func(i, j int) bool {
				return items[i].GetCreated().After(items[j].GetCreated())
			})
		}
	} else {
		if orderby == constant.OrderAsc {
			sort.Slice(items, func(i, j int) bool {
				return items[i].GetName() < items[j].GetName()
			})
		} else {
			sort.Slice(items, func(i, j int) bool {
				return items[i].GetName() > items[j].GetName()
			})
		}
	}
	return items
}

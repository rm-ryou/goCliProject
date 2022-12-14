package cmd

import (
	"example.com/GoCliProject/model"
)

type Sorter interface {
	Sort([]model.User)
}

type QuickSorter struct {
}

func quickSort(data []model.User, left, right int) {
	if right-left <= 1 {
		return
	}
	pivoIndex := (left + right) / 2
	pivo := data[pivoIndex].Id
	data[pivoIndex], data[right-1] = data[right-1], data[pivoIndex]
	i := left
	for j := left; j < right-1; j++ {
		if data[j].Id < pivo {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[right-1] = data[right-1], data[i]
	quickSort(data, left, i)
	quickSort(data, i+1, right)
}

func (s *QuickSorter) Sort(data []model.User) {
	quickSort(data, 0, len(data))
}

type MergeSorter struct {
}

func mergeSort(data []model.User, left, right int) {
	if right-left == 1 {
		return
	}
	pivoIndex := left + (right-left)/2
	mergeSort(data, left, pivoIndex)
	mergeSort(data, pivoIndex, right)
	var tmp []model.User
	for i := left; i < pivoIndex; i++ {
		tmp = append(tmp, data[i])
	}
	for i := right - 1; i >= pivoIndex; i-- {
		tmp = append(tmp, data[i])
	}
	leftIndex := 0
	rightIndex := len(tmp) - 1
	for i := left; i < right; i++ {
		if tmp[leftIndex].Id <= tmp[rightIndex].Id {
			data[i] = tmp[leftIndex]
			leftIndex++
		} else {
			data[i] = tmp[rightIndex]
			rightIndex--
		}
	}
}

func (s *MergeSorter) Sort(data []model.User) {
	mergeSort(data, 0, len(data))
}

type InsertSorter struct {
}

func insertSort(data []model.User, start, end int) {
	for i := start + 1; i < end; i++ {
		for j := i; j > start && data[j].Id < data[j-1].Id; j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
}

func (s *InsertSorter) Sort(data []model.User) {
	insertSort(data, 0, len(data))
}

func NewSorter(s string) Sorter {
	if s == "quick" {
		return new(QuickSorter)
	} else if s == "merge" {
		return new(MergeSorter)
	} else if s == "insertion" {
		return new(InsertSorter)
	} else {
		return nil
	}
}

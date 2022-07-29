package cmd

import (
	"testing"

	"example.com/GoCliProject/model"
)

var tests = []model.User{{74, "hoge", "hoge@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{94, "foo", "foo@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{42, "bar", "bar@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{24, "fuga", "fuga@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{59, "piyo", "piyo@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{-100, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{-1, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{0, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{2147483647, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{4242, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{424, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{1, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
}

var sorted = []model.User{{-100, "hoge", "hoge@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{-1, "foo", "foo@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{0, "bar", "bar@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{1, "fuga", "fuga@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{24, "piyo", "piyo@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{42, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{59, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{74, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{94, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{424, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{4242, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
	{2147483647, "xxxx", "xxxx@example.com", "xxxxxxxxxxx", "2022-01-01 00:00:00"},
}

func TestQuickSort(t *testing.T) {
	test := make([]model.User, len(tests))
	copy(test, tests)
	quickSort(test, 0, len(test))
	for i := 0; i < len(tests); i++ {
		if test[i].Id != sorted[i].Id {
			t.Errorf("sorted %v", sorted[i].Id)
			t.Errorf("test   %v", test[i].Id)
		}
	}
}

func TestMergeSort(t *testing.T) {
	test := make([]model.User, len(tests))
	copy(test, tests)
	mergeSort(test, 0, len(test))
	for i := 0; i < len(test); i++ {
		if test[i].Id != sorted[i].Id {
			t.Errorf("sorted %v", sorted[i].Id)
			t.Errorf("test   %v", test[i].Id)
		}
	}
}

func TestInsertSort(t *testing.T) {
	test := make([]model.User, len(tests))
	copy(test, tests)
	insertSort(test, 0, len(test))
	for i := 0; i < len(test); i++ {
		if test[i].Id != sorted[i].Id {
			t.Errorf("sorted %v", sorted[i].Id)
			t.Errorf("test   %v", test[i].Id)
		}
	}
}

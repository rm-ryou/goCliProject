package model

import (
	"encoding/json"
	"os"
)

type Args struct {
	SortType       string
	OutputJsonFile string
	Users          []User
}

func NewArgs(sortType, inputJsonFile, outputJsonFile string) (*Args, error) {
	var args = Args{
		SortType:       sortType,
		OutputJsonFile: outputJsonFile,
	}
	bytes, err := os.ReadFile(inputJsonFile)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &args.Users); err != nil {
		return nil, err
	}
	return &args, nil
}

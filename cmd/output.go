package cmd

import (
	"encoding/json"
	"os"

	"example.com/GoCliProject/model"
)

func OutputToFile(outputJsonFile string, userData []model.User) error {
	jsonData, err := json.MarshalIndent(userData, "", "\t")
	if err != nil {
		return err
	}
	file, err := os.Create(outputJsonFile)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(jsonData)
	return nil
}

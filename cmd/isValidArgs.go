package cmd

import (
	"fmt"
	_ "os"

	"example.com/GoCliProject/model"
)

type dataId struct {
	[]model.User.Id
}

func (id dataId) Execute() bool {
	return id <= 0
}

type dataName struct {
	[]model.User.Name
}

func (name dataName) Execute() bool {
	return name == ""
}

type dataEmail struct {
	[]model.User.Email
}

func (email dataEmail) Execute() bool {
	return email == ""
}

type dataPhoneNumber struct {
	[]model.User.PhoneNumber
}

func (phoneNumber dataPhoneNumber) Execute() bool {
	return phoneNumber == ""
}

type dataBirthday struct {
	[]model.User.Birthday
}

func (birthday dataBirthday) Execute() bool {
	return birthday == ""
}

func (data shape) Execute() bool {
	return data.Id >= 0
}

func Validater(data []model.User) bool {
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i].Name)
		data[i].Execute()
	}
	return true
}
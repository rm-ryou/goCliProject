package cmd

import (
	"example.com/GoCliProject/model"
	"fmt"
	_ "os"
)

func validId(data model.User) error {
	if data.Id <= 0 {
		return fmt.Errorf("invalid id")
	}
	return nil
}

func validName(data model.User) error {
	if data.Name == "" {
		return fmt.Errorf("invalid name")
	}
	return nil
}

func validEmail(data model.User) error {
	if data.Email == "" {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func validPhoneNumber(data model.User) error {
	if data.PhoneNumber == "" {
		return fmt.Errorf("invalid phone_number")
	}
	return nil
}

func validBirthday(data model.User) error {
	if data.Birthday == "" {
		return fmt.Errorf("invalid birthday")
	}
	return nil
}

func Validater(data []model.User) error {
	for i := 0; i < len(data); i++ {
		if err := validId(data[i]); err != nil {
			return err
		}
		if err := validName(data[i]); err != nil {
			return err
		}
		if err := validEmail(data[i]); err != nil {
			return err
		}
		if err := validPhoneNumber(data[i]); err != nil {
			return err
		}
		if err := validBirthday(data[i]); err != nil {
			return err
		}
	}
	return nil
}

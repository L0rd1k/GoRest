package models

import (
	"bookmarks-api/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"` //Пользователь, которому принадлежит этот контакт
}

func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return utils.Message(false, "Contact name should be on the payload"), false
	}
	if contact.Phone == "" {
		return utils.Message(false, "Phone number should be on the payload"), false
	}
	if contact.UserId <= 0 {
		return utils.Message(false, "User is not recognized"), false
	}
	//Все обязательные параметры присутствуют
	return utils.Message(true, "success"), true
}

func (contact *Contact) Create() map[string]interface{} {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}
	GetDataBase().Create(contact)
	resp := utils.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) *Contact {
	contact := &Contact{}
	err := GetDataBase().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {

	contacts := make([]*Contact, 0)
	err := GetDataBase().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}

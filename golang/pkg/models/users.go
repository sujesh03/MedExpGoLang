package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sujesh03/ExpenseTracker/pkg/config"
)

type Users struct {
	gorm.Model
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (u *Users) CreateUser() *Users {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []Users {
	var Users []Users
	db.Find(&Users)
	return Users
}

func DeleteUser(id int64) Users {
	var User Users
	db.Where("id = ?", id).Delete(&User)
	return User
}

func GetUserById(id int64) (*Users, *gorm.DB) {
	var User Users
	db.Where("id = ?", id).Find(&User)
	return &User, db
}

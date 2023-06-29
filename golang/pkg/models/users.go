package models

import (
	"fmt"

	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sujesh03/ExpenseTracker/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	gorm.Model
	Username string `json:"Username"`
	Password string `json:"-"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}
func GenerateHashedPassword(password string) (string, error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedpassword), nil
}
func (unc *Users) CreateUser() (*Users, error) {
	hashedPassword, _ := GenerateHashedPassword(unc.Password)
	unc.Password = hashedPassword

	// Check if the user already exists in the database
	var existingUser Users
	if err := db.Where("username = ?", unc.Username).First(&existingUser).Error; err == nil {
		// User with the same name already exists, return an error
		fmt.Println("User already exists")
		return nil, errors.New("User already exists")
	}

	if db.NewRecord(unc) {
		// Create the user record in the database
		if err := db.Create(unc).Error; err != nil {
			// Return an error if the creation process fails
			fmt.Println("Error in creating user", err)
			return nil, err

		}
	}

	// Return the user object and nil error on successful creation
	return unc, nil
}

func GetAllUsers() []Users {
	var Users []Users
	db.Find(&Users)
	return Users
}

func DeleteUser(id uint64) Users {
	var User Users
	db.Where("id = ?", id).Delete(&User)
	return User
}

func GetUserById(id uint64) (*Users, *gorm.DB) {
	var User Users
	db.Where("id = ?", id).Find(&User)
	return &User, db
}

package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sujesh03/ExpenseTracker/pkg/config"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponceModel struct {
	AccessToken string   `json:"AccessToken"`
	UserName    string   `json:"UserName"`
	Error       AppError `json:"error"`
}

type Claims struct {
	Username string `json:"username"`
	UserId   uint   `json:"userid"`
	jwt.StandardClaims
}

var jwtKey = []byte("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")

func init() {
	config.Connect()
	db = config.GetDB()
	// db.AutoMigrate(&Skills{})
}

func GetUserFromName(name string) (*Users, error) {
	var getUser Users
	db := db.Where("username=?", name).Find(&getUser)
	if db.Error != nil {
		return nil, db.Error
	}
	return &getUser, nil
}

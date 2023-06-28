package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sujesh03/ExpenseTracker/pkg/config"
)

type Expenses struct {
	gorm.Model
	UserId      int
	Date        time.Time `json:"Date"`
	ProjectCode string    `json:"ProjectCode"`
	VendorName  string    `json:"VendorName"`
	Description string    `json:"Description"`
	Type        string    `json:"Type"`
	Amount      float64   `json:"Amount"`
	Status      int       `json:"Status"`
	Users       Users     `gorm:"foreignkey:UserId;references:ID;" json:"UserId"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (e *Expenses) CreateExpense() *Expenses {
	user := &Users{}
	// fmt.Println("User ID: ", e.UserId)
	// fmt.Println("Status: ", e.Status)
	if err := db.Where("ID=?", e.UserId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User not found, error: ", err)
			return nil
		}
		return nil
	}
	e.Users = *user
	db.NewRecord(e)
	db.Create(&e)
	return e
}

func GetAllExpenses() []Expenses {
	var Expenses []Expenses
	db.Preload("Users").Find(&Expenses)
	return Expenses
}

func DeleteExpense(id int64) Expenses {
	var Expense Expenses
	db.Where("id = ?", id).Delete(&Expense)
	return Expense
}

func GetExpenseById(id int64) (*Expenses, *gorm.DB) {
	var Expense Expenses
	db.Where("id = ?", id).Find(&Expense)
	return &Expense, db
}

func GetExpenseUserById(Id uint) *Users {
	var User Users
	db.Where("ID=?", Id).Find(&User)
	return &User
}

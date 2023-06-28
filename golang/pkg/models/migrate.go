package models

import "github.com/sujesh03/ExpenseTracker/pkg/config"

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Users{})
	db.AutoMigrate(&Vendors{})
	db.AutoMigrate(&Expenses{})

	db.Model(&Expenses{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&Expenses{}).Association("Users")
	db.Model(&Vendors{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&Vendors{}).Association("Users")
}

package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sujesh03/ExpenseTracker/pkg/config"
)

var db *gorm.DB

type Vendors struct {
	gorm.Model
	UserId   uint64    `json:"UserId"`
	Name     string    `json:"Name"`
	Status   int       `json:"Status"`
	Created  time.Time `json:"Created"`
	Modified time.Time `json:"Modified"`
	Users    Users     `gorm:"foreignkey:UserId;references:ID;"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func (v *Vendors) CreateVendor() *Vendors {
	user := &Users{}
	if err := db.Where("ID=?", v.UserId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return nil
	}
	v.Users = *user
	db.NewRecord(v)
	db.Create(&v)
	return v
}

func GetAllVendors() []Vendors {
	var Vendors []Vendors
	db.Preload("Users").Find(&Vendors)
	return Vendors
}

func DeleteVendor(id uint64) Vendors {
	var Vendor Vendors
	db.Where("ID=?", id).Delete(&Vendor)
	return Vendor
}

func GetVendorById(id uint64) (*Vendors, *gorm.DB) {
	var Vendor Vendors
	db.Where("ID=?", id).Find(&Vendor)
	return &Vendor, db
}

func GetVendorUserById(Id uint64) *Users {
	var User Users
	db.Where("ID = ?", Id).Find(&User)
	return &User
}

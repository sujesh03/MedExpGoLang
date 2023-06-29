package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sujesh03/ExpenseTracker/pkg/models"
	"github.com/sujesh03/ExpenseTracker/pkg/utils"
)

func CreateVendor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	CreateVendor := &models.Vendors{}
	utils.ParseBody(r, CreateVendor)
	e := CreateVendor.CreateVendor()
	res, _ := json.Marshal(e)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetVendors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	vendors := models.GetAllVendors()
	res, _ := json.Marshal(vendors)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteVendor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	vars := mux.Vars(r)
	vendorid := vars["id"]
	ID, _ := strconv.ParseUint(vendorid, 0, 0)
	vendor := models.DeleteVendor(ID)
	res, _ := json.Marshal(vendor)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateVendor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	vendor := &models.Vendors{}
	utils.ParseBody(r, vendor)
	vars := mux.Vars(r)
	vendorid := vars["id"]
	ID, _ := strconv.ParseUint(vendorid, 0, 0)
	vendordetails, db := models.GetVendorById(ID)
	if vendor.UserId != 0 {
		if vendor.UserId != vendordetails.UserId {
			user := models.GetVendorUserById((vendor.UserId))
			vendordetails.UserId = vendor.UserId
			vendordetails.Users = *user
		}
	}
	if vendor.Name != "" {
		vendordetails.Name = vendor.Name
	}
	if vendor.Status != 0 {
		vendordetails.Status = vendor.Status
	}
	if vendor.Created.IsZero() == false {
		vendordetails.Created = vendor.Created
	}
	if vendor.Modified.IsZero() == false {
		vendordetails.Modified = vendor.Modified
	}
	db.Save(&vendordetails)
	res, _ := json.Marshal(vendordetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

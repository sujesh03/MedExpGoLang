package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sujesh03/ExpenseTracker/pkg/models"
	"github.com/sujesh03/ExpenseTracker/pkg/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	CreateUser := &models.Users{}
	utils.ParseBody(r, CreateUser)
	e := CreateUser.CreateUser()
	res, _ := json.Marshal(e)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	users := models.GetAllUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	vars := mux.Vars(r)
	userid := vars["id"]
	ID, _ := strconv.ParseInt(userid, 0, 0)
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	user := &models.Users{}
	utils.ParseBody(r, user)
	vars := mux.Vars(r)
	userid := vars["id"]
	ID, _ := strconv.ParseInt(userid, 0, 0)
	userdetails, db := models.GetUserById(ID)
	if user.Username != "" {
		userdetails.Username = user.Username
	}
	if user.Password != "" {
		userdetails.Password = user.Password
	}
	db.Save(&userdetails)
	res, _ := json.Marshal(userdetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

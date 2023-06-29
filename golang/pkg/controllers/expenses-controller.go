package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sujesh03/ExpenseTracker/pkg/models"
	"github.com/sujesh03/ExpenseTracker/pkg/utils"
)

func CreateExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	Createexpenses := &models.Expenses{}
	utils.ParseBody(r, Createexpenses)
	// fmt.Printf("request", r.Body)
	e := Createexpenses.CreateExpense()
	res, _ := json.Marshal(e)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	expenses := models.GetAllExpenses()
	res, _ := json.Marshal(expenses)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	vars := mux.Vars(r)
	expensesid := vars["id"]
	ID, _ := strconv.ParseUint(expensesid, 0, 0)
	expense := models.DeleteExpense(ID)
	res, _ := json.Marshal(expense)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	expense := &models.Expenses{}
	utils.ParseBody(r, expense)
	vars := mux.Vars(r)
	expensesid := vars["id"]
	ID, _ := strconv.ParseUint(expensesid, 0, 0)
	expensedetails, db := models.GetExpenseById(ID)
	if expense.UserId != 0 {
		if expensedetails.UserId != expense.UserId {
			user := models.GetExpenseUserById((expense.UserId))
			expensedetails.UserId = expense.UserId
			expensedetails.Users = *user
		}
	}
	if expense.ProjectCode != "" {
		expensedetails.ProjectCode = expense.ProjectCode
	}
	if expense.VendorName != "" {
		expensedetails.VendorName = expense.VendorName
	}
	if expense.Description != "" {
		expensedetails.Description = expense.Description
	}
	if expense.Type != "" {
		expensedetails.Type = expense.Type
	}
	if expense.Amount != 0 {
		expensedetails.Amount = expense.Amount
	}
	if expense.Status != 0 {
		expensedetails.Status = expense.Status
	}
	db.Save(&expensedetails)
	res, _ := json.Marshal(expensedetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

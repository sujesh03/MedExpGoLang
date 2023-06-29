package routes

import (
	"github.com/gorilla/mux"
	"github.com/sujesh03/ExpenseTracker/pkg/controllers"
)

var RegisterETStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/api/v1/login/", controllers.Login).Methods("POST")
	r.HandleFunc("/api/v1/expenses/", controllers.CreateExpenses).Methods("POST")
	r.HandleFunc("/api/v1/expenses/", controllers.GetExpenses).Methods("GET")
	r.HandleFunc("/api/v1/expenses/{expensesid}", controllers.UpdateExpense).Methods("PUT")
	r.HandleFunc("/api/v1/expenses/{expensesid}", controllers.DeleteExpense).Methods("DELETE")

	r.HandleFunc("/api/v1/vendors/", controllers.CreateVendor).Methods("POST")
	r.HandleFunc("/api/v1/vendors/", controllers.GetVendors).Methods("GET")
	r.HandleFunc("/api/v1/vendors/{vendorid}", controllers.UpdateVendor).Methods("PUT")
	r.HandleFunc("/api/v1/vendors/{vendorid}", controllers.DeleteVendor).Methods("DELETE")

	r.HandleFunc("/api/v1/users/", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users/", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/users/{userid}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/v1/users/{userid}", controllers.DeleteUser).Methods("DELETE")
}

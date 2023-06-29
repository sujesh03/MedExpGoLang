package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sujesh03/ExpenseTracker/pkg/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	jwtKey          = []byte("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")
	ErrInvalidInput = models.AppError{Code: http.StatusBadRequest, Message: "Invalid request payload"}
	ErrInternal     = models.AppError{Code: http.StatusInternalServerError, Message: "Internal server error"}
	ErrUnauthorized = models.AppError{Code: http.StatusUnauthorized, Message: "Invalid credentials"}
	ErrUserNotFound = models.AppError{Code: http.StatusUnauthorized, Message: "User not found"}
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		writeErrorResponse(w, ErrInvalidInput)
		return
	}

	EmpCredential, err := models.GetUserFromName(creds.Username)
	if err != nil {
		writeErrorResponse(w, ErrUserNotFound)
		return
	}

	fmt.Println("name : ", EmpCredential.Username)
	fmt.Println("password : ", EmpCredential.Password)
	fmt.Println("creds.Password : ", creds.Password)

	if checkPasswordHash(creds.Password, EmpCredential.Password) {
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &models.Claims{
			Username: creds.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			writeErrorResponse(w, ErrInternal)
			return
		}

		responce := &models.LoginResponceModel{}
		responce.AccessToken = tokenString
		responce.Error.Error = false
		jsonResponse, _ := json.Marshal(responce)

		if err != nil {
			writeErrorResponse(w, ErrInternal)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	} else {
		writeErrorResponse(w, ErrUnauthorized)
		return
	}

}

func writeErrorResponse(w http.ResponseWriter, err models.AppError) {
	w.WriteHeader(err.Code)
	responce := &models.LoginResponceModel{}
	responce.AccessToken = ""
	responce.Error = err
	responce.Error.Error = true
	//user.Error.ErrorData = err
	jsonResponse, _ := json.Marshal(responce)
	//jsonResponse, _ := json.Marshal(err)
	w.Write(jsonResponse)
}

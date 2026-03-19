package handlers

import (
	"encoding/json"
	"leave-management/database"
	"leave-management/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request){
	var login models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil{
		http.Error(w,"Invalid requst",http.StatusBadRequest)
		return
	}

	var id int
	var hashedPassword string
	var roleID int

	err = database.DB.QueryRow(
		`SELECT id,password,role_id 
		FROM employees
		WHERE email=$1`,
		login.Email,
	).Scan(&id,&hashedPassword,&roleID)

	if err != nil{
		http.Error(w,"Invalid Email",http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(login.Password),
	)

	if err != nil{
		http.Error(w,"Invalid password",http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"employee_id":id,
		"role_id":roleID,
	})

}
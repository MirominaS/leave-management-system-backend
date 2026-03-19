package handlers

import (
	"encoding/json"
	"leave-management/database"
	"leave-management/models"
	"net/http"
)


func GetEmployees(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
	SELECT e.id,e.name,e.email,r.role_name
	FROM employees e
	JOIN roles r ON e.role_id = r.id
	`)

	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var employees []map[string]interface{}

	for rows.Next(){
		var id int
		var name string
		var email string
		var role string

		rows.Scan(&id,&name,&email,&role)

		emp := map[string]interface{}{
			"id": id,
			"name": name,
			"email": email,
			"role": role,
		}
		employees = append(employees, emp)
	}
	json.NewEncoder(w).Encode(employees)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request){
	var emp models.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	query := `INSERT INTO employees (name,email,role_id)
	VALUES ($1,$2,$3) RETURNING id`

	err = database.DB.QueryRow(query,emp.Name,emp.Email,emp.RoleID).Scan(&emp.ID)

	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(emp)
}
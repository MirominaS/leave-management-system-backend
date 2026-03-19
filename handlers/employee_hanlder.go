package handlers

import (
	"encoding/json"
	"leave-management/database"
	"leave-management/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

	hashedPassword,err := bcrypt.GenerateFromPassword(
		[]byte(emp.Password),
		bcrypt.DefaultCost,
	)

	if err != nil{
		http.Error(w,"Error hashing Password",http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO employees (name,email,password,role_id)
	VALUES ($1,$2,$3,$4) RETURNING id`

	err = database.DB.QueryRow(query,emp.Name,emp.Email,string(hashedPassword),emp.RoleID).Scan(&emp.ID)

	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message":"Employee created successfully",
	})
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
    empID := r.URL.Query().Get("id")
    var emp models.Employee
    err := database.DB.QueryRow("SELECT id, name, email FROM employees WHERE id = $1", empID).Scan(&emp.ID, &emp.Name, &emp.Email)
    
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(emp)
}
package handlers

import (
	"encoding/json"
	"leave-management/database"
	"leave-management/models"
	"net/http"
)


func CreateLeave(w http.ResponseWriter, r *http.Request){
	var leave models.LeaveRequest

	err := json.NewDecoder(r.Body).Decode(&leave)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	var pendingCount int

	err = database.DB.QueryRow(
		"SELECT Count(*) FROM leave_request WHERE employee_id=$1 AND status_id=1",
		leave.EmployeeID,
	).Scan(&pendingCount)

	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	if pendingCount > 0 {
		http.Error(w, "Employee already has a pending leave",http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(
		`INSERT INTO leave_request(employee_id,start_date,end_date,leave_type_id,reason,status_id)
		VALUES ($1,$2,$3,$4,$5,1)`,		
		leave.EmployeeID,
		leave.StartDate,
		leave.EndDate,
		leave.LeaveTypeID,
		leave.Reason,
	)

	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Leave request submitted successfully",
	})
}

func GetLeaves(w http.ResponseWriter, r*http.Request){

	rows, err := database.DB.Query(`
	SELECT 
	l.id,
	l.employee_id,
	e.name,
	e.email,
	lt.leave_type_name,
	l.start_date,
	l.end_date,
	l.reason,
	ls.status_name
	FROM leave_request l
	JOIN employees e ON l.employee_id = e.id
	JOIN leave_types lt ON l.leave_type_id = lt.id
	JOIN leave_status ls ON l.status_id = ls.id`)

	if err != nil{
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var leaves []map[string]interface{}
	for rows.Next(){
		var id int
		var employeeID int
		var name string
		var email string
		var start string
		var end string
		var leaveType string
		var reason string
		var status string

		rows.Scan(&id,&employeeID,&name,&email,&leaveType,&start,&end,&reason,&status)

		leave := map[string]interface{}{
			"id": id,
			"employee_id": employeeID,
			"employee": name,
			"email":email,
			"leave_type_name":leaveType,
			"start_date": start,
			"end_date": end,
			"reason": reason,
			"status": status,
		}
		leaves = append(leaves, leave)
	}
	
	json.NewEncoder(w).Encode(leaves)
}

func GetLeaveTypes(w http.ResponseWriter, r *http.Request){
	rows, err := database.DB.Query("SELECT id, leave_type_name FROM leave_types")
	if err != nil{
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var types []map[string]interface{}
	for rows.Next(){
		var id int
		var name string
		rows.Scan(&id,&name)
		types = append(types, map[string]interface{}{
			"id":id,
			"name":name,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((types))
}

func ApproveLeave(w http.ResponseWriter,r *http.Request){
	role := r.Header.Get("role_id")

	if role != "4" && role != "6"{
		http.Error(w,"Only Manager or Admin can approve leave",http.StatusForbidden)
		return
	}

	id := r.URL.Query().Get("id")

	_,err := database.DB.Exec(
		"UPDATE leave_request SET status_id = 2 WHERE id=$1",
		id,
	)

	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message":"Leave approved",
	})
}

func RejectLeave(w http.ResponseWriter, r *http.Request){
	role := r.Header.Get("role_id")

	if role != "4" && role != "6"{
		http.Error(w,"Only Manager or Admin can reject leave",http.StatusForbidden)
		return
	} 

	id := r.URL.Query().Get("id")

	_, err := database.DB.Exec(
		"UPDATE leave_request SET status_id = 3 WHERE id=$1",
		id,
	)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message":"Leave rejected",
	})
}

func CancelLeave(w http.ResponseWriter, r *http.Request){

	employeeID := r.Header.Get("employee_id")
	leaveID := r.URL.Query().Get("id")

	_,err := database.DB.Exec(
		`DELETE FROM leave_request WHERE id=$1
		AND employee_id=$2`,
		leaveID,employeeID,
	)

	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message":"Leave cancelled",
	})
}
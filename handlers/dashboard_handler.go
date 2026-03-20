package handlers

import (
	"encoding/json"
	"leave-management/database"
	"leave-management/models"
	"net/http"
)

func GetDashboard(w http.ResponseWriter, r *http.Request){
	var dashboard models.Dashboard

	database.DB.QueryRow(
		"SELECT COUNT(*) FROM employees",
	).Scan(&dashboard.TotalEmployees)

	database.DB.QueryRow(
		"SELECT COUNT(*) FROM leave_request WHERE status_id = 1",
	).Scan(&dashboard.PendingLeaves)

	database.DB.QueryRow(
		"SELECT COUNT(*) FROM leave_request WHERE status_id = 2",
	).Scan(&dashboard.ApprovedLeaves)

	database.DB.QueryRow(
		"SELECT COUNT(*) FROM leave_request WHERE status_id = 3",
	).Scan(&dashboard.RejectedLeaves)

	rows,err := database.DB.Query(`
	SELECT r.id, r.role_name,
    	COUNT(e.id) FILTER (WHERE l.id IS NULL) as active_count,
    	COUNT(e.id) FILTER (WHERE l.id IS NOT NULL) as away_count
	FROM roles r
	LEFT JOIN employees e ON r.id = e.role_id
	LEFT JOIN leave_request l ON e.id = l.employee_id
    	AND l.status_id = 2
    	AND CURRENT_DATE BETWEEN l.start_date AND l.end_date
	GROUP BY r.id, r.role_name`)

	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next(){
		var rs models.RoleStatus 
		if err := rows.Scan(&rs.ID,&rs.Title,&rs.Active,&rs.OnLeave); err != nil{
			continue
		}
		dashboard.Roles = append(dashboard.Roles, rs)
	}
	
	json.NewEncoder(w).Encode(dashboard)
}

func GetRecentActivity(w http.ResponseWriter, r *http.Request){
	rows, err := database.DB.Query(`
		SELECT lr.id, e.name, ls.status_name, lr.created_at 
		FROM leave_request lr
		JOIN employees e ON lr.employee_id = e.id
		JOIN leave_status ls ON lr.status_id = ls.id
		ORDER BY lr.created_at DESC
		LIMIT 5
	`)

	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return 
	}
	defer rows.Close()

	var activities []models.Activity

	for rows.Next(){
		var a models.Activity
		rows.Scan(&a.ID,&a.Name,&a.Status,&a.CreatedAt)
		activities = append(activities,a)
	}

	json.NewEncoder(w).Encode(activities)
}
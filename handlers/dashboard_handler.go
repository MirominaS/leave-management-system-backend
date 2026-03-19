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
	
	json.NewEncoder(w).Encode(dashboard)
}
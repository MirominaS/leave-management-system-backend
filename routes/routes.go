package routes

import (
	"leave-management/handlers"
	"net/http"
)

func RegisterRoutes() *http.ServeMux{
	mux := http.NewServeMux()

	mux.HandleFunc("/login",handlers.Login)

	mux.HandleFunc("/employees", handlers.GetEmployees)
	mux.HandleFunc("/employees/create", handlers.CreateEmployee)

	mux.HandleFunc("/leave", handlers.GetLeaves)
	mux.HandleFunc("/leave/create", handlers.CreateLeave)
	mux.HandleFunc("/leave/cancel",handlers.CancelLeave)
	mux.HandleFunc("/leave/approve",handlers.ApproveLeave)
	mux.HandleFunc("/leave/reject",handlers.RejectLeave)


	mux.HandleFunc("/dashboard", handlers.GetDashboard)
	
	return mux
}
package models

type LeaveRequest struct{
	ID int `json:"id"`
	EmployeeID int `json:"employee_id"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	LeaveTypeID int `json:"leave_type_id"`
	Reason string `json:"reason"`
	StatusID int `json:"status_id"`
}
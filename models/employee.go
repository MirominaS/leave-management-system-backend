package models

type Employee struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	RoleID int `json:"role_id"`
}
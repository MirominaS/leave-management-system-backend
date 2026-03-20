package models

type RoleStatus struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Active int `json:"active"`
	OnLeave int `json:"onLeave"`
}
type Activity struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
}
type Dashboard struct {
	TotalEmployees int `json:"total_employees"`
	PendingLeaves  int `json:"pending_leaves"`
	ApprovedLeaves int `json:"approved_leaves"`
	RejectedLeaves int `json:"rejected_leaves"`
	Roles []RoleStatus `json:"roles"`
}
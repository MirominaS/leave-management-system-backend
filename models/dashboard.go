package models

type Dashboard struct {
	TotalEmployees int `json:"total_employees"`
	PendingLeaves  int `json:"pending_leaves"`
	ApprovedLeaves int `json:"approved_leaves"`
	RejectedLeaves int `json:"rejected_leaves"`
}
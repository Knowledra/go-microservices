package models

type UpdateDepartment struct {
	Name string `json:"name" binding:"omitempty"`
	Code string `json:"code" binding:"omitempty"`
}

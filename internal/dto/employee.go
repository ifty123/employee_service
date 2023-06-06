package dto

type EmployeeResponse struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

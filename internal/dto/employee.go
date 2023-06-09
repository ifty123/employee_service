package dto

type EmployeeResponse struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type EmployeeWithJWTResponse struct {
	EmployeeResponse
	JWT string `json:"jwt"`
}

package dto

type EmployeeResponse struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type EmployeeWithJWTResponse struct {
	EmployeeResponse
	JWT string `json:"jwt"`
}

package model

import "latihan_service/internal/dto"

type Employee struct {
	Common
	Fullname   string `json:"fullname" gorm:"varchar;not_null"`
	Email      string `json:"email" gorm:"varchar;not_null;unique"`
	Password   string `json:"password" gorm:"varchar;not_null"`
	DivisionId uint   `json:"division_id"`
	Role       string `json:"role"`
}

func (emp *Employee) UpdateEmployee(payload *dto.UpdateEmployeeReq) {
	if payload.Fullname != "" || payload.Fullname != emp.Fullname {
		emp.Fullname = payload.Fullname
	}

	if payload.Email != "" || payload.Email != emp.Email {
		emp.Email = payload.Email
	}

	if payload.Role != "" || payload.Role != emp.Role {
		emp.Role = payload.Role
	}

}

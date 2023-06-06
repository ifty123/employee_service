package model

type Employee struct {
	Common
	Fullname string `json:"fullname" gorm:"varchar;not_null"`
	Email    string `json:"email" gorm:"varchar;not_null;unique"`
	Password string `json:"password" gorm:"varchar;not_null"`
}

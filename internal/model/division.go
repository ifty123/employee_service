package model

type Division struct {
	NameDivision string `json:"name_division" gorm:"varchar;not_null"`
	CodeDivision string `json:"code_division" gorm:"varchar;not_null;unique"`
	Common
}

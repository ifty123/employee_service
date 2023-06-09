package seeder

import (
	"latihan_service/database"

	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeeder() *seed {
	return &seed{database.GetConnection()}
}

func (s *seed) SeedAll() {
	EmployeeSeeder(s.DB)
	divisionSeeder(s.DB)
}

func (s *seed) DeleteAll() {
	s.DB.Exec("DELETE FROM employees")
	s.DB.Exec("DELETE FROM divisions")
}

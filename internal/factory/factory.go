package factory

import (
	"latihan_service/database"
	"latihan_service/internal/repository"
)

type Factory struct {
	EmployeeRepository repository.EmployeeRepository
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewEmployeeRepository(db),
	}
}

package seeder

import (
	"latihan_service/internal/model"
	"log"

	"gorm.io/gorm"
)

func EmployeeSeeder(db *gorm.DB) {
	var emp = []model.Employee{
		{
			Fullname: "Alif",
			Email:    "alifipa5@gmail.com",
			Role:     "Employee",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
		},
		{
			Fullname: "Qotrun",
			Email:    "qotrunnada5@gmail.com",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
		},
	}

	if err := db.Create(&emp).Error; err != nil {
		log.Printf("cannot seed data employees, with error %v\n", err)
	}
	log.Println("success seed data employees")
}

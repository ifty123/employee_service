package seeder

import (
	"latihan_service/internal/model"
	"log"
	"time"

	"gorm.io/gorm"
)

func divisionSeeder(db *gorm.DB) {
	now := time.Now()
	var division = []model.Division{
		{
			NameDivision: "Animation",
			CodeDivision: "AN",
			Common: model.Common{
				ID:        1,
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	}

	//create
	if err := db.Create(&division).Error; err != nil {
		log.Printf("cannot seed data divisions, with error %v\n", err)
	}
	log.Println("success seed data divisions")
}

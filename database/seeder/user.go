package seeder

import (
	"log"
	"time"

	"submission-5/internal/model"

	"gorm.io/gorm"
)

func userSeeder(db *gorm.DB) {
	now := time.Now()
	var employees = []model.User{
		{
			Fullname: "Hikayat",
			Email:    "hikayat@gmail.com",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			RoleID:   1,
			Common:   model.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
		},
		{
			Fullname: "Hanindito",
			Email:    "hanindito@gmail.com",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			RoleID:   2,
			Common:   model.Common{ID: 2, CreatedAt: now, UpdatedAt: now},
		},
		{
			Fullname: "vincentlhubbard",
			Email:    "vincentlhubbard@gmail.com",
			Password: "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			RoleID:   2,
			Common:   model.Common{ID: 3, CreatedAt: now, UpdatedAt: now},
		},
	}
	if err := db.Create(&employees).Error; err != nil {
		log.Printf("cannot seed data employees, with error %v\n", err)
	}
	log.Println("success seed data employees")
}

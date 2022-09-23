package seeder

import (
	"github.com/hikayatz/go-deploy-heroku/database"
	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeeder() *seed {
	return &seed{database.GetConnection()}
}

func (s *seed) SeedAll() {
	roleSeeder(s.DB)
	userSeeder(s.DB)
}

func (s *seed) DeleteAll() {
	s.DB.Exec("DELETE FROM users")
	s.DB.Exec("DELETE FROM roles")
	s.DB.Exec("DELETE FROM books")
}

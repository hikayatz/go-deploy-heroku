package factory

import (
	"github.com/hikayatz/go-deploy-heroku/database"
	"github.com/hikayatz/go-deploy-heroku/internal/repository"
)

type Factory struct {
	UserRepository repository.User
	RoleRepository repository.Role
	BookRepository repository.Book
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewUserRepository(db),
		repository.NewRoleRepository(db),
		repository.NewBookRepository(db),
	}
}

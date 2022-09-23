package factory

import (
	"submission-5/database"
	"submission-5/internal/repository"
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

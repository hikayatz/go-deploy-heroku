package auth

import (
	"context"
	"strings"
	"testing"

	"submission-5/database"
	"submission-5/database/seeder"
	"submission-5/internal/dto"
	"submission-5/internal/factory"

	"github.com/stretchr/testify/assert"
)

func TestAuthServiceLoginByEmailAndPasswordSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		payload     = dto.ByEmailAndPasswordRequest{
			Email:    "vincentlhubbard@gmail.com",
			Password: "123abcABC!",
		}
	)
	res, err := authService.LoginByEmailAndPassword(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(payload.Email, res.Email)
	asserts.Len(strings.Split(res.JWT, "."), 3)
}

func TestAuthServiceLoginByEmailAndPasswordRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	var (
		asserts     = assert.New(t)
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		payload     = dto.ByEmailAndPasswordRequest{
			Email:    "azkaframadhan@gmail.com",
			Password: "123abcABC!",
		}
	)
	_, err := authService.LoginByEmailAndPassword(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestAuthServiceLoginByEmailAndPasswordunmatchedEmailAndPassword(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		payload     = dto.ByEmailAndPasswordRequest{
			Email:    "vincentlhubbard@gmail.com",
			Password: "1234567890",
		}
	)
	_, err := authService.LoginByEmailAndPassword(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 400")
	}
}

func TestAuthServiceRegisterByEmailAndPasswordSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		payload     = dto.RegisterUserRequestBody{
			Fullname: "Azka Fadhli Ramadhan",
			Email:    "azkaframadhan@gmail.com",
			Password: "123abcABC!",
		}
	)
	payload.FillDefaults()
	res, err := authService.RegisterByEmailAndPassword(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
	asserts.Equal(payload.Fullname, res.Fullname)
	asserts.Equal(payload.Email, res.Email)
	asserts.Len(strings.Split(res.JWT, "."), 3)
}

func TestAuthServiceRegisterByEmailAndPasswordUserExist(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		payload     = dto.RegisterUserRequestBody{
			Fullname: "Vincent L Hubbard",
			Email:    "vincentlhubbard@gmail.com",
			Password: "123abcABC!",
		}
	)
	_, err := authService.RegisterByEmailAndPassword(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 409")
	}
}

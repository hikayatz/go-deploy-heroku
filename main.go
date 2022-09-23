package main

import (
	"flag"
	"log"
	"os"

	"github.com/hikayatz/go-deploy-heroku/database"
	"github.com/hikayatz/go-deploy-heroku/database/migration"
	"github.com/hikayatz/go-deploy-heroku/database/seeder"
	"github.com/hikayatz/go-deploy-heroku/internal/factory"
	"github.com/hikayatz/go-deploy-heroku/internal/http"
	"github.com/hikayatz/go-deploy-heroku/internal/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
	database.GetConnection()
}

func main() {
	database.CreateConnection()

	var m string // for check migration
	var s string // for check seeder

	flag.StringVar(
		&m,
		"m",
		"none",
		`this argument for check if user want to migrate table, rollback table, or status migration
to use this flag:
	use -m=migrate for migrate table
	use -m=rollback for rollback table
	use -m=status for get status migration`,
	)

	flag.StringVar(
		&s,
		"s",
		"none",
		`this argument for check if user want to seed table to use this flag: use -s=all to seed all table`,
	)

	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
	} else if m == "rollback" {
		migration.Rollback()
	} else if m == "status" {
		migration.Status()
	}
	migration.Migrate()
	if s == "all" {
		seeder.NewSeeder().DeleteAll()
		seeder.NewSeeder().SeedAll()
	}

	f := factory.NewFactory()
	e := echo.New()

	middleware.LogMiddlewares(e)

	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

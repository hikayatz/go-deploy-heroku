package database

import (
	"sync"

	"github.com/hikayatz/go-deploy-heroku/pkg/util"
	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

func CreateConnection() {
	conf := dbConfig{
		User: util.Getenv("DB_USER", "root"),
		Pass: util.Getenv("DB_PASS", "1234567890"),
		Host: util.Getenv("DB_HOST", "localhost"),
		Port: util.Getenv("DB_PORT", "3306"),
		Name: util.Getenv("DB_NAME", "alterra_agmc"),
	}

	mysql := mysqlConfig{dbConfig: conf}
	once.Do(func() {
		mysql.Connect()
	})
}

func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}

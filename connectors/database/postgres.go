package database

import (
	"fmt"
	"os"

	"github.com/itispx/eruka/aws/secretsmanager"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectPostgresDB() error {
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbSecretArn := os.Getenv("DB_SECRET_ARN")
	sslmode := "disable"

	if os.Getenv("ENV") != "dev" {
		credentials, err := secretsmanager.GetDBCredentials(&dbSecretArn)
		if err != nil {
			return err
		}

		dbName = credentials.DBName
		port = fmt.Sprintf(`%v`, credentials.Port)
		user = credentials.Username
		password = credentials.Password
		sslmode = "require"
	}

	dsn := fmt.Sprintf(`
	 host=%s 
	 user=%s 
	 password=%s 
	 dbname=%s 
	 port=%v 
	 sslmode=%s`,
		host,
		user,
		password,
		dbName,
		port,
		sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	DB = db
	return nil
}

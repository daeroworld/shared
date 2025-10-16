package database

import (
	"fmt"

	"github.com/daeroworld/shared/configuration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresqlWrapper struct {
	Driver *gorm.DB
}

func ConnectPostgresqlDatabase(database *configuration.Database) *PostgresqlWrapper {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=whiscript port=9920 sslmode=disable TimeZone=Asia/Seoul", database.Uri, database.Username, database.Password)
	driver, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return nil
	}
	return &PostgresqlWrapper{
		Driver: driver,
	}
}

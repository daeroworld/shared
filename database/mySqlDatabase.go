package database

import (
	"fmt"
	"log"
	"mntreamer/shared/configuration"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlWrapper struct {
	Driver *gorm.DB
}

func ConnectMysqlDatabase(database *configuration.Database) *MysqlWrapper {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/mntreamer?charset=utf8mb4&parseTime=True&loc=Local", database.Username, database.Password, database.Uri)
	driver, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return &MysqlWrapper{
		Driver: driver,
	}
}

package storage

import (
	"fmt"

	"tglanz/gorm-example/model"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseParams struct {
	User     string
	Password string
	Host     string
	Port     uint
	Database string
}

func CreateMysqlDatabase(params DatabaseParams) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		params.User, params.Password, params.Host, params.Port, params.Database)

	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		panic("unable to connect to the database")
	}

	return db
}

func CreatePostgresDatabase(params DatabaseParams) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		params.Host, params.User, params.Password, params.Database, params.Port)

	dialector := postgres.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		panic("unable to connect to the database")
	}

	return db
}

func AutoMigrate(db *gorm.DB) {
	var command model.Command
	var spec model.Spec
	var commandResult model.CommandResult
	var context model.Context

	// db.AutoMigrate(&spec, &command, &context, &commandResult)
	db.AutoMigrate(&spec)
	db.AutoMigrate(&command)
	db.AutoMigrate(&context)
	db.AutoMigrate(&commandResult)
}

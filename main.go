package main

import (
	"fmt"
	"tglanz/gorm-example/storage"
)

func main() {
	fmt.Println("Creating datbase context")
	db := storage.CreatePostgresDatabase(storage.DatabaseParams{
		User:     "local",
		Password: "local",
		Host:     "localhost",
		Port:     5432, //3306,
		Database: "local",
	})

	fmt.Println("Performing migration")
	storage.AutoMigrate(db)

}

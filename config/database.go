package config

import (
	sqlite "github.com/glebarez/sqlite"
	gorm "gorm.io/gorm"
	schema "gorm.io/gorm/schema"

	entities "github.com/eolme/backmemes/entities"
	seed "github.com/eolme/backmemes/seed"
)

var Database *gorm.DB

func migrate() error {
	return Database.AutoMigrate(
		&entities.Author{},
		&entities.Meme{},
	)
}

func Connect() error {
	var err error

	// включаем общий кэш sqlite для всех процессов
	Database, err = gorm.Open(sqlite.Open("runtime/database.sqlite?cache=shared&_pragma=foreign_keys(1)"), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})

	if err != nil {
		panic(err)
	}

	migrate()
	seed.Seed(Database)

	return nil
}

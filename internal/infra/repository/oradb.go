package repository

import (
	"fmt"

	cmnMod "ng_autocode_back_service/common/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbCore struct {
	Db *gorm.DB
}

type databaseSet struct {
	ReadDB  dbCore
	WriteDB dbCore
}

func ReadEbatestConnect() (*gorm.DB, error) {
	// Connect to the database
	return connect(cmnMod.READ_DB)
}
func WriteEbatestConnect() (*gorm.DB, error) {
	// Connect to the database
	return connect(cmnMod.WRITE_DB)
}

func connect(ReadWrite int8) (*gorm.DB, error) {
	dsn := "user=postgres password=admin dbname=Finservice host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("Connected to database successfully, mode:", ReadWrite)
	return db, nil
}

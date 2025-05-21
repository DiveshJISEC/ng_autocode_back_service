package repositoy

import (
	"fmt"

	cmnMod "ng_autocode_back_service/common/model"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

type dbCore struct {
	Db *gorm.DB
}

type daatbaseSet struct {
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
	dsn := "oracle://ebatest:ebatest10.57.72.47:2227/ebatest"
	// Connect to the database
	db, err := gorm.Open(oracle.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	schemas := fmt.Sprintf("ALTER SESSION SET CURRENT_SCHEMA %= s", "ebatest")
	if err := db.Exec(schemas).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	fmt.Println("Connected to database")
	return db, nil
}

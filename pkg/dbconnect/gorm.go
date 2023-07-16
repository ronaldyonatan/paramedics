package dbconnect

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGORM(dbConfig DBConfig) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s ",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Dbuser,
		dbConfig.Dbpassword,
		dbConfig.Dbname,
		dbConfig.Sslmode,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbPing, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = dbPing.Ping()
	if err != nil {
		return nil, err
	}

	return
}

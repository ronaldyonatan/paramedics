package dbconnect

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectSqlx(dbConfig DBConfig) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Dbuser,
		dbConfig.Dbpassword,
		dbConfig.Dbname,
		dbConfig.Sslmode,
	)
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return
}

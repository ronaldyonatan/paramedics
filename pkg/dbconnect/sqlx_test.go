package dbconnect

import (
	"testing"
)

func TestConnectSqlx(t *testing.T) {

	db, err := ConnectSqlx(DBConfig{
		Host:       "172.16.20.40",
		Port:       "5432",
		Dbname:     "hacktiv8Nando",
		Dbuser:     "sa",
		Dbpassword: "P@ssw0rd",
		Sslmode:    "disable",
	})
	if err != nil {
		t.Fatalf("Error connect to database: %v", err)
	}
	if db == nil {
		t.Fatal("Error connect to database: nil db")
	}
}

package users

import (
	"testing"

	"github.com/fernandojec/assignment-2/pkg/dbconnect"
	"github.com/fernandojec/assignment-2/pkg/utils"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func Test_authRepo_InsertAuth(t *testing.T) {
	_ = godotenv.Load("../../.env")
	dbx, _ := dbconnect.ConnectSqlx(dbconnect.DBConfig{
		Host:       utils.GetEnv("POSTGRES_HOST"),
		Port:       utils.GetEnv("POSTGRES_PORT"),
		Dbname:     utils.GetEnv("POSTGRES_DBNAME"),
		Dbuser:     utils.GetEnv("POSTGRES_USER"),
		Dbpassword: utils.GetEnv("POSTGRES_PASSWORD"),
		Sslmode:    utils.GetEnv("POSTGRES_SSLMODE"),
	})
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		data auth
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  uint
		wantErr bool
	}{
		{
			name: "Insert Auth",
			fields: fields{
				db: dbx,
			},
			args: args{
				data: auth{
					Email:     "fernando.riyo@jec.co.id",
					Password:  "123123",
					Username:  "fernando",
					FirstName: "Fernando",
					LastName:  "Riyo",
				},
			},
			wantErr: false,
			wantId:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := authRepo{
				db: tt.fields.db,
			}
			gotId, err := r.InsertAuth(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("authRepo.InsertAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("authRepo.InsertAuth() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

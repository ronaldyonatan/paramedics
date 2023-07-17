package users

import (
	"github.com/fernandojec/assignment-2/pkg/utils"
	"github.com/jmoiron/sqlx"
)

type authRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) authRepo {
	return authRepo{db}
}

func (r authRepo) InsertAuth(data auth) (id uint, err error) {
	query := "Insert into auths (email,password,username,first_name,last_name,is_active,created_at) values ($1,$2,$3,$4,$5,$6,$7) returning id"

	stmn, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	err = stmn.QueryRow(
		utils.NewSQLNullString(data.Email),
		utils.NewSQLNullString(data.Password),
		utils.NewSQLNullString(data.Username),
		utils.NewSQLNullString(data.FirstName),
		utils.NewSQLNullString(data.LastName),
		data.IsActive,
		data.CreatedAt,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return
}

func (r authRepo) InsertVerifyAuth(data verifyAuth) (err error) {
	query := `INSERT INTO verify_auths(
		auth_id, token, created_at, expired_at)
		VALUES ($1,$2,$3,$4);`

	stmn, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmn.Exec(
		data.AuthId,
		utils.NewSQLNullString(data.Token),
		data.CreatedAt,
		data.ExpiredAt,
	)
	if err != nil {
		return err
	}
	return
}
func (r authRepo) GetAuthByID(id uint) (data auth, err error) {
	query := `select * from auths where id=$1`

	_, err = r.db.Prepare(query)
	if err != nil {
		return auth{}, err
	}
	err = r.db.Get(&data, query, id)

	return
}
func (r authRepo) GetAuthByEmail(email string) (data auth, err error) {
	query := `select * from auths where email=$1`

	_, err = r.db.Prepare(query)
	if err != nil {
		return auth{}, err
	}
	err = r.db.Get(&data, query, email)

	return
}
func (r authRepo) UpdateAuthIsActive(data auth) (err error) {
	query := `UPDATE auths
	SET activated_at=$1,is_active=$2 where id=$3;
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.ActivatedAt.Time, data.IsActive, data.Id)
	if err != nil {
		return err
	}
	return
}
func (r authRepo) GetVerifyAuthByToken(token string) (data verifyAuth, err error) {
	query := `select * from verify_auths where token = $1`

	_, err = r.db.Prepare(query)
	if err != nil {
		return verifyAuth{}, err
	}
	err = r.db.Get(&data, query, token)

	return
}

func (r authRepo) UpdateVerifyAuth(data verifyAuth) (err error) {
	query := `UPDATE verify_auths
	SET activated_at=$1 where id=$2;
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.ActivatedAt.Time, data.Id)
	if err != nil {
		return err
	}
	return
}

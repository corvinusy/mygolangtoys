package models

import (
	"database/sql"
	"errors"
	"time"
)

type User struct {
	UserId              int
	UserLogin           string
	UserName            string
	UserFio             sql.NullString
	UserUserdata        []byte // JSON
	UserRegDate         time.Time
	UserMasterDeviceId  int
	UserBillRec         int
	UserRecoveryPhone   sql.NullString
	UserRecoveryEmail   sql.NullString
	UserIsActive        bool
	UserLastContactTime time.Time
}

/*----------------------------------------------------------------------------*/
func (ct *Controller) CreateUser(db *sql.DB, login string, name string, hash []byte) error {
	var (
		uid int
		err error
	)
	// insert user
	err = db.QueryRow("SELECT msc.insert_new_user2($1::text, $2::text, $3::text)",
		login, name, hash).Scan(&uid)
	switch {
	case err == sql.ErrNoRows:
		return errors.New("database connection error")
	case err != nil:
		return err
	case err == nil && uid == 0:
		return errors.New("User insert error")
	}

	return nil
}

/*----------------------------------------------------------------------------*/
func (u *User) GetUser(db *sql.DB, uid int) error {

	// get user from database
	err := db.QueryRow("SELECT * FROM msc.get_user_row2($1::int)", uid).Scan(
		&u.UserId,
		&u.UserLogin,
		&u.UserName,
		&u.UserFio,
		&u.UserUserdata,
		&u.UserRegDate,
		&u.UserMasterDeviceId,
		&u.UserBillRec,
		&u.UserRecoveryPhone,
		&u.UserRecoveryEmail,
		&u.UserIsActive,
		&u.UserLastContactTime)
	switch {
	case err == sql.ErrNoRows:
		return errors.New("user not found")
	case err != nil:
		return err
	}

	return nil
}

/*----------------------------------------------------------------------------*/
func (u *User) GetUserAuth(db *sql.DB, ulogin string) (string, error) {
	var passHash sql.NullString

	// get user from database
	err := db.QueryRow("SELECT * FROM msc.get_user_auth_row2($1::int)", uid).Scan(
		&u.UserId,
		&u.UserLogin,
		&u.UserName,
		&u.UserFio,
		&u.UserUserdata,
		&u.UserRegDate,
		&u.UserMasterDeviceId,
		&u.UserBillRec,
		&u.UserRecoveryPhone,
		&u.UserRecoveryEmail,
		&u.UserIsActive,
		&u.UserLastContactTime,
		&passHash)
	switch {
	case err == sql.ErrNoRows:
		return "", errors.New("user not found")
	case err != nil:
		return "", err
	}

	return passHash, nil

}

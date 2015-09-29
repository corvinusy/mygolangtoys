package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang/glog"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId              int
	UserLogin           string
	UserName            string
	UserFio             string
	UserUserdata        []byte // JSON
	UserRegDate         time.Time
	UserMasterDeviceId  int
	UserBillRec         int
	UserRecoveryPhone   string
	UserRecoveryEmail   string
	UserIsActive        bool
	UserLastContactTime time.Time
}

func (user *User) HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		glog.Fatalf("Couldn't hash password: %v", err)
		return nil, err
	}
	return hash, nil
}

func GetUserByEmail(database *sql.DB, email string) (user *User) {
	return nil
}

func (u *User) Create(db *sql.DB, hash []byte) error {
	err := db.QueryRow("SELECT msc.insert_new_user2($1::text, $2::text, $3::text)",
		u.UserLogin, u.UserName, hash).Scan(&u.UserId)

	if err != nil || u.UserId == 0 {
		return errors.New("User insert error")
	}

	return nil
}

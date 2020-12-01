package dao

import (
	"database/sql"
	"week02/util"

	"github.com/pkg/errors"
)

// ErrRecordNotFound ...
var ErrRecordNotFound = errors.New("record not found")

// User ...
type User struct {
	ID string `json:"id"`
}

// One 查询一条
func One(user User) (User, error) {
	err := sql.ErrNoRows
	if errors.Is(err, sql.ErrNoRows) {
		return user, errors.WithStack(ErrRecordNotFound)
	}
	if err != nil {
		return user, nil
	}
	return user, nil
}

// Rows 查询多条
func Rows(users []*User) (*util.PageRes, error) {
	err := sql.ErrNoRows
	res := &util.PageRes{}
	res.Total = 0
	if errors.Is(err, sql.ErrNoRows) {
		res.Rows = users
		return res, nil
	}
	if err != nil {
		return res, nil
	}
	return res, nil
}

func (u *User) Error() string {
	return ErrRecordNotFound.Error()
}

package service

import (
	"week02/dao"

	xerrors "github.com/pkg/errors"
)

//ServiceQuerysql ...
func ServiceQuerysql(s string) (string, error) {
	v, err := dao.MysqlQuery(s)
	if err != nil {
		return "", xerrors.Wrap(err, "Querying key ["+s+"] failed.")
	}
	return v, nil
}

func ServiceAddsql(k string, v string) error {
	return xerrors.Wrap(dao.MysqlAdd(k, v), "Adding key "+k+" failed.")
}

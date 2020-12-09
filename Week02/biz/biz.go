package biz

import (
	dao "Week02/dao"
	"database/sql"

	"github.com/pkg/errors"
)

func Biz(mode string) (*dao.RtnTables, error) {
	var name string = "gogogo"
	usertable, err := dao.Dao(mode)
	if err != nil {
		//add biz level error handle, not throw out the root error
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("0 row finded.")
		}
		return nil, errors.WithMessage(err, "biz() error with input: "+name)
	}
	return usertable, err
}

package biz

import(
	dao "Week02/dao"
	"github.com/pkg/errors"
)

func Biz(mode string) (*dao.RtnTables, error) {
	var name string = "gogogo"
	usertable,err := dao.Dao(mode)
	if err!= nil{
		return nil, errors.WithMessage(err, "biz() error with input: " + name)
	}
	return usertable, err
}

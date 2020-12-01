package dao

import (
	sql "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type RtnTables struct{
	ID   string
	name string
}

func dbSearchTable() (table *RtnTables,err error){
	return nil, sql.ErrNoRows
}

func dbSearchTable1() (table *RtnTables,err error){
	return nil, errors.New("unknown errors")
}

func Dao(input string) (table *RtnTables, err error){
	//db,err:= sql.Open("mysql","root:123456@tcp(10.253.48.53:3306)/gogogo")
	//defer db.close()
	//Here comes out the err
	if input == "ErrNoRows"{
		table , err = dbSearchTable()
	} else {
		table , err = dbSearchTable1()
	}
	
	switch{
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, errors.Wrap(err, "dao error")
	default:
		return table, err
	}
}

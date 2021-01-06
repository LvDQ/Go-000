package main

import (
	"errors"
	"fmt"
)

// dao
func BatchGetGirl() ([]Girl, error) {
	rows, err := db.Query("SELECT * FROM girls WHERE love =10")
	if err != nil {

	}
	err := rows.Err()
	if err != nil {
		//opqure sql.ErrNoRows
		//bussines code
		//stack trace
		return errors, Wrapf(code.ErrNotFound, fmt.Sprintf("query: %s failed(%v)", sql, err))
	}
	return []Girl{}, nil
}

//biz
func Usecase() error {
	v, err := BatchGetGirl()
	if errors.Is(code.ErrNotFOund, err) {

	}
}

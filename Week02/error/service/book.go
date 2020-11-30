package service

import (
	"database/sql"
	"error/dao"
	"fmt"

	"github.com/pkg/errors"
)

// GetBookNameByID service: get book name by book id
func GetBookNameByID(id int) (string, error) {
	book, err := dao.GetBookByID(id)
	if errors.Is(err, sql.ErrNoRows) {
		return "", nil
	}
	if err != nil {
		err = errors.WithStack(err)
		return "", err
	}
	fmt.Println(book.ID)
	return book.Name, nil
}

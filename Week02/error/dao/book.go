package dao

import (
	"database/sql"
	"error/model"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// GetBookByID get all book model
func GetBookByID(id int) (*model.Book, error) {
	// ${username}: must override your username
	// ${passwerd}: must override your password
	// ${database}: must override your database
	db, err := sql.Open("mysql", "${username}:${passwerd}@/${database}")
	if err != nil {
		fmt.Println("error 1")
		return nil, err
	}
	defer db.Close()

	var book model.Book
	row := db.QueryRow("select * from book where id=?", id)
	if err := row.Scan(&book.ID, &book.Name, &book.Intro); err != nil {
		return nil, err
	}
	fmt.Printf("id:%d,name:%s,intro:%s", book.ID, book.Name, book.Intro)
	return &book, nil
}

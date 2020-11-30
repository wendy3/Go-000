package main

import (
	"error/api"
	"fmt"

	"github.com/pkg/errors"
)

func fn() error {
	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")
	return errors.Wrap(e3, "outer")
}

func main() {
	fmt.Println("hello world")

	// learn pkg/errors
	err := fn()
	fmt.Println(err)
	fmt.Println(errors.Cause(err))

	// define and init param(id)
	id := 1
	// // call api to test error demo
	name, err := api.GetBookByID(id)
	if err != nil {
		fmt.Println(errors.Cause(err))
		return
	}
	fmt.Printf("book name: %s", name)
}

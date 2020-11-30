package api

import "error/service"

// GetBookByID service written by wendy
func GetBookByID(id int) (string, error) {
	return service.GetBookNameByID(id)
}

package utils

import (
	"strconv"
)

func ValidateID(id string) (int, bool) {
	var valid = true
	productID, err := strconv.Atoi(id)
	if err != nil {
		valid = false
		return 0, valid
	}

	return productID, valid
}

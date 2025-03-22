package dto

import (
	"github.com/go-playground/validator/v10"
)

type ProductDTO struct {
	ID    int
	Name  string  `validate:"required,min=3,max=255"`
	Price float64 `validate:"required,gt=0"`
}

func (v ProductDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(v)
}

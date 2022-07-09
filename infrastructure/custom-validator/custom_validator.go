package customvalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type (
	CustomValidator interface {
		UniqueValidator(fl validator.FieldLevel) bool
		ExistValidator(fl validator.FieldLevel) bool
	}
	customValidator struct {
		db *sqlx.DB
	}
)

func New(db *sqlx.DB) CustomValidator {
	return &customValidator{
		db: db,
	}
}

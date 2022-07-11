package user

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/michaelputeraw/krobot-auth-service/model/database"
)

func (r *repository) buildInsertQuery(input *database.User) sq.InsertBuilder {
	vals := sq.Eq{
		"id":         input.ID,
		"full_name":  input.FullName,
		"gender":     input.Gender,
		"email":      input.Email,
		"password":   input.Password,
		"created_at": time.Now().Unix(),
		"updated_at": time.Now().Unix(),
	}
	insertBuilder := sq.Insert(r.GetTableName()).SetMap(vals)
	return insertBuilder
}

func (r *repository) buildSelectQuery() sq.SelectBuilder {
	selection := []string{
		"id",
		"full_name",
		"gender",
		"email",
		"password",
		"created_at",
		"updated_at",
		"deleted_at",
	}
	selectBuilder := sq.Select(selection...).Where(sq.Eq{"deleted_at": nil}).From(r.GetTableName())
	return selectBuilder
}

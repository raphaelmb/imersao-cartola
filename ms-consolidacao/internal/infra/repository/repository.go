package repository

import (
	"errors"

	"github.com/raphaelmb/imersao-cartola-consolidacao/internal/infra/db"
)

var ErrQueriesNotSet = errors.New("queries not set")

type Repository struct {
	*db.Queries
}

func (r *Repository) SetQuery(q *db.Queries) {
	r.Queries = q
}

func (r *Repository) Validade() error {
	if r.Queries == nil {
		return ErrQueriesNotSet
	}
	return nil
}

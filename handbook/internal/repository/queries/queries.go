package queries

import (
	"github.com/mayamika/mai-lev/handbook/internal/postgres"
)

type Queries struct {
	pool *postgres.Pool
}

func New(pool *postgres.Pool) *Queries {
	return &Queries{
		pool: pool,
	}
}

package repository

import (
	"context"

	"github.com/mayamika/mai-lev/handbook/internal/model"
	"github.com/mayamika/mai-lev/handbook/internal/postgres"
	"github.com/mayamika/mai-lev/handbook/internal/repository/queries"
)

type repo struct {
	*queries.Queries
}

type Repository interface {
	Entries(context.Context) ([]model.Entry, error)
	AddEntry(context.Context, model.Entry) error
	UpdateEntry(context.Context, string, model.Entry) error
	DeleteEntry(context.Context, string) error
}

func NewRepository(pool *postgres.Pool) Repository {
	return &repo{
		Queries: queries.New(pool),
	}
}

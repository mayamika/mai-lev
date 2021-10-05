package service

import (
	"context"

	"github.com/mayamika/mai-lev/handbook/internal/model"
)

type EntriesRepo interface {
	Entries(ctx context.Context) ([]model.Entry, error)
	AddEntry(ctx context.Context, entry model.Entry) error
	UpdateEntry(ctx context.Context, id string, entry model.Entry) error
	DeleteEntry(ctx context.Context, id string) error
}

type EntriesService struct {
	repo EntriesRepo
}

func NewEntriesService(repo EntriesRepo) *EntriesService {
	return &EntriesService{
		repo: repo,
	}
}

func (es *EntriesService) Entries(ctx context.Context) ([]model.Entry, error) {
	return es.repo.Entries(ctx)
}

func (es *EntriesService) AddEntry(ctx context.Context, entry model.Entry) error {
	return es.repo.AddEntry(ctx, entry)
}

func (es *EntriesService) UpdateEntry(
	ctx context.Context,
	id string,
	entry model.Entry,
) error {
	return es.repo.UpdateEntry(ctx, id, entry)
}

func (es *EntriesService) DeleteEntry(
	ctx context.Context,
	id string,
) error {
	return es.repo.DeleteEntry(ctx, id)
}

package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/mayamika/mai-lev/handbook/internal/model"
)

type EntriesService interface {
	Entries(ctx context.Context) ([]model.Entry, error)
	AddEntry(ctx context.Context, entry model.Entry) error
	UpdateEntry(ctx context.Context, id string, entry model.Entry) error
	DeleteEntry(ctx context.Context, id string) error
}

type Entries struct {
	service EntriesService
}

func NewEntries(service EntriesService) *Entries {
	return &Entries{
		service: service,
	}
}

func (e *Entries) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route(`/`, func(r chi.Router) {
		r.Get(`/`, e.entries)
		r.Post(`/`, e.addEntry)
		r.Patch(`/{id}`, e.updateEntry)
		r.Delete(`/{id}`, e.deleteEntry)
	})

	return r
}

func (e *Entries) entries(w http.ResponseWriter, r *http.Request) {
	entries, err := e.service.Entries(r.Context())
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.PlainText(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, entries)
}

func (e *Entries) addEntry(w http.ResponseWriter, r *http.Request) {
	var entry model.Entry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}
	if err := e.service.AddEntry(r.Context(), entry); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, entry.ID)
}

func (e *Entries) updateEntry(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, `id`)

	var entry model.Entry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}
	if err := e.service.UpdateEntry(r.Context(), id, entry); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, entry.ID)
}

func (e *Entries) deleteEntry(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, `id`)

	if err := e.service.DeleteEntry(r.Context(), id); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, id)
}

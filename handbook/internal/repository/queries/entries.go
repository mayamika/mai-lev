package queries

import (
	"context"
	"fmt"

	"github.com/mayamika/mai-lev/handbook/internal/model"
)

const selectEntriesQuery = `SELECT id, name, surname, university FROM entries`

func (q *Queries) Entries(ctx context.Context) ([]model.Entry, error) {
	rows, err := q.pool.Query(ctx, selectEntriesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []model.Entry
	for rows.Next() {
		var e model.Entry
		scanErr := rows.Scan(
			&e.ID,
			&e.Name,
			&e.Surname,
			&e.University,
		)
		if scanErr != nil {
			return nil, scanErr
		}
		entries = append(entries, e)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

const insertEntryQuery = `INSERT INTO entries(id, name, surname, university)
						VALUES($1, $2, $3, $4)`

func (q *Queries) AddEntry(
	ctx context.Context,
	entry model.Entry,
) error {
	_, insertErr := q.pool.Exec(ctx, insertEntryQuery,
		entry.ID,
		entry.Name,
		entry.Surname,
		entry.University,
	)
	return insertErr
}

const updateEntryQuery = `UPDATE entries SET id = $2, name = $3, surname = $4, university = $5
						WHERE id = $1`

func (q *Queries) UpdateEntry(
	ctx context.Context,
	id string,
	entry model.Entry,
) error {
	_, updateErr := q.pool.Exec(ctx, updateEntryQuery,
		id,
		entry.ID,
		entry.Name,
		entry.Surname,
		entry.University,
	)
	return updateErr
}

const deleteEntryQuery = `DELETE FROM entries WHERE id = $1`

func (q *Queries) DeleteEntry(ctx context.Context, id string) error {
	_, deleteErr := q.pool.Exec(ctx, deleteEntryQuery, id)
	fmt.Println(id)
	return deleteErr
}

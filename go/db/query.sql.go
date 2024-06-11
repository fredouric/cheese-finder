// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
)

const addCheese = `-- name: AddCheese :exec
INSERT INTO cheese (departement, fromage, pagefrancaise, englishpage, lait, geoshape, geopoint2d) VALUES (?, ?, ?, ?, ?, ?, ?)
`

type AddCheeseParams struct {
	Departement   string
	Fromage       string
	Pagefrancaise string
	Englishpage   string
	Lait          string
	Geoshape      string
	Geopoint2d    string
}

func (q *Queries) AddCheese(ctx context.Context, arg AddCheeseParams) error {
	_, err := q.db.ExecContext(ctx, addCheese,
		arg.Departement,
		arg.Fromage,
		arg.Pagefrancaise,
		arg.Englishpage,
		arg.Lait,
		arg.Geoshape,
		arg.Geopoint2d,
	)
	return err
}

const deleteAllCheeses = `-- name: DeleteAllCheeses :exec
DELETE FROM cheese
`

func (q *Queries) DeleteAllCheeses(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllCheeses)
	return err
}

const getAllCheeses = `-- name: GetAllCheeses :many
SELECT id, departement, fromage, pagefrancaise, englishpage, lait, geoshape, geopoint2d FROM cheese
ORDER BY fromage
`

func (q *Queries) GetAllCheeses(ctx context.Context) ([]Cheese, error) {
	rows, err := q.db.QueryContext(ctx, getAllCheeses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Cheese
	for rows.Next() {
		var i Cheese
		if err := rows.Scan(
			&i.ID,
			&i.Departement,
			&i.Fromage,
			&i.Pagefrancaise,
			&i.Englishpage,
			&i.Lait,
			&i.Geoshape,
			&i.Geopoint2d,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCheese = `-- name: GetCheese :one
SELECT id, departement, fromage, pagefrancaise, englishpage, lait, geoshape, geopoint2d FROM cheese 
WHERE id=? LIMIT 1
`

func (q *Queries) GetCheese(ctx context.Context, id int64) (Cheese, error) {
	row := q.db.QueryRowContext(ctx, getCheese, id)
	var i Cheese
	err := row.Scan(
		&i.ID,
		&i.Departement,
		&i.Fromage,
		&i.Pagefrancaise,
		&i.Englishpage,
		&i.Lait,
		&i.Geoshape,
		&i.Geopoint2d,
	)
	return i, err
}
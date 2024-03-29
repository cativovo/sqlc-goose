// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getTodos = `-- name: GetTodos :many
SELECT id, task, completed, created_at, description FROM todo
`

func (q *Queries) GetTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.Query(ctx, getTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Task,
			&i.Completed,
			&i.CreatedAt,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertTodo = `-- name: InsertTodo :one
INSERT INTO todo (task, description, completed, created_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING id, task, completed, created_at, description
`

type InsertTodoParams struct {
	Task        string
	Description pgtype.Text
	Completed   pgtype.Bool
}

func (q *Queries) InsertTodo(ctx context.Context, arg InsertTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, insertTodo, arg.Task, arg.Description, arg.Completed)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Task,
		&i.Completed,
		&i.CreatedAt,
		&i.Description,
	)
	return i, err
}

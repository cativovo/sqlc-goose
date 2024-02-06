-- name: GetTodos :many
SELECT * FROM todo;

-- name: InsertTodo :one
INSERT INTO todo (task, description, completed, created_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING *;

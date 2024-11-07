-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY id;

-- name: CreateTask :one
INSERT INTO tasks (due, status, priority, title, description) VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET due = $2, status = $3, priority = $4, title = $5, description = $6
WHERE id = $1
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;

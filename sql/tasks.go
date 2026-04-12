package sql

import (
	"context"
	"study/task"

	"github.com/jackc/pgx/v5"
)

func CreateTask(conn pgx.Conn, ctx context.Context, task *task.Task) error {
	queryStr := `
	INSERT INTO tasks (title, description, completed, time_created)
	VALUES ($1,$2,$3,$4);`

	_, err := conn.Exec(ctx, queryStr,
		task.Title,
		task.Description,
		task.IsCompleted,
		task.TimeCreated)
	return err
}

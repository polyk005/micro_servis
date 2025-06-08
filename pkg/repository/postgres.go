package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	query := `
	INSERT INTO tasks (id, type, status, params, created_at)
	VALUES (:id, :type, :status, :params, :created_at)
	`
	task.ID = uuid.New().String()
	task.Status = "pending"
	task.CreatedAt = time.Now()
	_, err := r.db.NamedExecContext(ctx, query, task)
	return err
}

func (r *PostgresRepository) UpdateTaskStatus(
	ctx context.Context,
	taskID string,
	status string,
	result string,
) error {
	query := `
	        UPDATE tasks
			SET status = $1, result = $2, comleted_at = NOW()
			WHERE id = $3
	`
	_, err := r.db.ExecContext(ctx, query, status, result, taskID)
	return err
}

func (r *PostgresRepository) GetTask(ctx context.Context, taskID string) (models.Task, error) {
	var task models.Task
	query := `SELECT * FROM tasks WHERE id = $1`
	err := r.db.GetContext(ctx, &task, query, taskID)
	if err == sql.ErrNoRows {
		  	return nil, nil
	}
	return &task, err
}
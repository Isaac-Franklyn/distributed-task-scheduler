package databases

import (
	"context"
	"fmt"
	"log"

	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"
	"github.com/jackc/pgx/v5"
)

type TaskDb struct {
	conn *pgx.Conn
}

func StartNewCockroachDb(connString string) (*TaskDb, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to CockroachDB: %w", err)
	}
	return &TaskDb{conn: conn}, nil
}

func (db *TaskDb) Close() {
	db.conn.Close(context.Background())
}

func (db *TaskDb) CreateTaskTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		type STRING NOT NULL,
		payload JSONB NOT NULL,
		priority INT NOT NULL,
		status STRING NOT NULL,
		retries INT NOT NULL
	);
	`
	_, err := db.conn.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to create tasks table: %w", err)
	}
	log.Println("Task table is ready")
	return nil
}

func (db *TaskDb) SaveTaskToDb(task *models.Task) error {

	query := `
	INSERT INTO tasks (id, type, payload, priority, status, retries)
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := db.conn.Exec(context.Background(), query,
		task.ID, task.Type, task.Payload, task.Priority, task.Status, task.Retries)
	if err != nil {
		return fmt.Errorf("failed to insert task: %w", err)
	}
	return nil
}

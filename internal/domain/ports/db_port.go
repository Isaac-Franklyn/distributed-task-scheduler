package ports

import "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"

type DbService interface {
	Close()
	SaveTaskToDb(task *models.Task) error
	CreateTaskTable()
}

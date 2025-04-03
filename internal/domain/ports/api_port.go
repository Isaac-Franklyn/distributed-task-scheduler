package ports

import model "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"

type APIService interface {
	Validate(task *model.Task) error
}

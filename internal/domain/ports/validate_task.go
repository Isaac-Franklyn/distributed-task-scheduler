package ports

import model "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"

type ValidateTask interface {
	Validate(task *model.Task) error
}

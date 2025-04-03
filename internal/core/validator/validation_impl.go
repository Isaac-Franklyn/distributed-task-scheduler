package api

import (
	"fmt"

	model "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"
	ports "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/ports"
)

type TaskValidatorImpl struct{}

func NewTaskValidatorImpl() ports.ValidateTask {
	return &TaskValidatorImpl{}
}

func (s *TaskValidatorImpl) Validate(task *model.Task) error {

	if task.Payload == nil || task.Payload == "" {
		return fmt.Errorf("payload is empty")
	}

	if task.Type != "api_call" {
		return fmt.Errorf("wrong task type, expected: 'api_call'")
	}

	if task.Priority <= 0 || task.Priority > 10 {
		return fmt.Errorf("priority is out of bounds, allowed: 1-10")
	}

	if task.Retries < 0 || task.Retries > 5 {
		return fmt.Errorf("retry limit exceeded, allowed: 0-5")
	}

	return nil
}

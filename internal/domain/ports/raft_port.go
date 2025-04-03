package ports

import "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"

type RaftService interface {
	StartCluster(n int)
	GetLeader() (*models.Node, error)
	SendTaskToCluster(task *models.Task) error
}

package api_test

import (
	"testing"

	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/core/api"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"
	"github.com/stretchr/testify/assert"
)

func TestApi_InvalidTasks(t *testing.T) {
	api := api.NewApiValidator()

	tests := []struct {
		name     string
		task     models.Task
		expected string
	}{
		{"empty payload", models.Task{Payload: nil, Type: "api_call", Priority: 5, Retries: 2}, "payload is empty"},
		{"invalid priority", models.Task{Payload: map[string]interface{}{"url": "http://example.com"}, Type: "api_call", Priority: 0, Retries: 2}, "priority is out of bounds, allowed: 1-10"},
		{"invalid task type", models.Task{Payload: map[string]interface{}{"url": "http://example.com"}, Type: "wrong_type", Priority: 5, Retries: 2}, "wrong task type, expected: 'api_call'"},
		{"invalid retries", models.Task{Payload: map[string]interface{}{"url": "http://example.com"}, Type: "api_call", Priority: 5, Retries: 8}, "retry limit exceeded, allowed: 0-5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := api.Validate(&tt.task)
			assert.Error(t, err)
			assert.Equal(t, tt.expected, err.Error())
		})
	}
}

func TestApi_ValidTask(t *testing.T) {
	api := api.NewApiValidator()

	task := &models.Task{
		Payload:  map[string]interface{}{"url": "www.example.com"},
		Type:     "api_call",
		Priority: 5,
		Retries:  3,
	}

	t.Run("valid task", func(t *testing.T) {
		err := api.Validate(task)
		assert.NoError(t, err)
	})
}

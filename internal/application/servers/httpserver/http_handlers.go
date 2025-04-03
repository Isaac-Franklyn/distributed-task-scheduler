package servers

import (
	"encoding/json"
	"net/http"

	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostTask(api ports.APIService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var task = &models.Task{}

		decoder := json.NewDecoder(ctx.Request.Body)
		if err := decoder.Decode(task); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := api.Validate(task); err != nil {
			ctx.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
			return
		}

		task.ID = uuid.NewString()
		task.Retries++
		task.Status = "Pending"

		ctx.JSON(http.StatusOK, gin.H{"message": "Task received", "task": task})
	}
}

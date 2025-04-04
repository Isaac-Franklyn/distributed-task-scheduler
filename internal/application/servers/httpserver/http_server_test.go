package httpserver_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/application/servers/httpserver"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type Response struct {
	Message string      `json:"message"`
	Task    models.Task `json:"task"`
}

func TestHttpServer(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockcluster := mocks.NewMockRaftService(ctrl)
	mockapi := mocks.NewMockAPIService(ctrl)

	mockcluster.EXPECT().CommitTaskToCluster(gomock.Any()).Return(nil)
	mockapi.EXPECT().Validate(gomock.Any()).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/task", httpserver.PostTask(mockapi, mockcluster))

	taskJSON := `{"type":"api_call", "payload": {"url":"https://example.com"}, "priority":1, "retries": 4}`
	req := httptest.NewRequest("POST", "/task", strings.NewReader(taskJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var res Response
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Now, res.Task contains the actual task struct
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 5, res.Task.Retries)
	assert.Equal(t, "Pending", res.Task.Status)
	assert.NotEmpty(t, res.Task.ID)

}

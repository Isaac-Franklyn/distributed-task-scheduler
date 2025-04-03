package httpserver

import (
	"log"

	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/ports"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	engine *gin.Engine
	api    ports.APIService
}

func NewHTTPServer(api ports.APIService) *HTTPServer {

	server := &HTTPServer{
		engine: gin.Default(),
		api:    api,
	}

	server.SetupRoutes()

	return server

}

func (srv *HTTPServer) SetupRoutes() {
	srv.engine.POST("/submit-task", PostTask(srv.api))
}

func (srv *HTTPServer) Start() error {
	log.Println("Starting HTTP server on :8080")
	return srv.engine.Run(":8080")
}

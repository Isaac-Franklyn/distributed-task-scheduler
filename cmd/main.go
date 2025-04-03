package main

import (
	servers "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/application/servers/httpserver"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/core/api"
)

func main() {

	api := api.NewApiValidator()

	srv := servers.NewHTTPServer(api)
	srv.Start()

}

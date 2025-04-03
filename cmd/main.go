package main

import (
	servers "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/application/servers/httpserver"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/core/api"
)

func main() {

	//setup the dependencies
	api := api.NewApiValidator()

	//start the external agents, servers, dbs, raft, etc
	srv := servers.NewHTTPServer(api)
	srv.Start()

}

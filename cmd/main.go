package main

import (
	servers "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/application/servers/httpserver"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/core/api"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/core/raft"
)

func main() {

	//setup the dependencies
	cluster := raft.NewRaftCluster()
	cluster.StartCluster(5)
	api := api.NewApiValidator()

	//start the external agents, servers, dbs, raft, etc
	srv := servers.NewHTTPServer(api, cluster)
	srv.Start()

}

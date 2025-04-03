package main

import (
	"log"

	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/application/databases"
	servers "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/application/servers/httpserver"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/core/api"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/core/raft"
)

func main() {

	//setup the dependencies
	cluster := raft.NewRaftCluster()
	cluster.StartCluster(5)

	connString := "postgresql://root@localhost:26257/distributed_task_scheduler?sslmode=disable"
	db, err := databases.StartNewCockroachDb(connString)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	api := api.NewApiValidator()

	//start the external agents, servers, dbs, raft, etc
	srv := servers.NewHTTPServer(api, cluster)
	srv.Start()

}

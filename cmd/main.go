package main

import (
	"log"
	"os"

	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/application/databases"
	servers "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/application/servers/httpserver"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/core/api"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/core/raft"
)

func main() {

	//setup the dependencies
	cluster := raft.NewRaftCluster()
	cluster.StartCluster(1)

	connString := os.Getenv("DATABASE_URL")
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

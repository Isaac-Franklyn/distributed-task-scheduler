package raft

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

type RaftCluster struct {
	Cluster []*models.Node
}

func NewRaftCluster() *RaftCluster {
	return &RaftCluster{}
}

func (raftcluster *RaftCluster) StartCluster(n int) {

	for i := 0; i < n; i++ {

		id := fmt.Sprintf("node-%d", i+1)
		addr := fmt.Sprintf("127.0.0.1.%d", 9000+i)
		node := createRaftNode(id, addr)
		raftcluster.Cluster = append(raftcluster.Cluster, node)
	}
}

func (raftcluster *RaftCluster) GetLeader() (*models.Node, error) {

	for _, node := range raftcluster.Cluster {
		if node.Raft.State() == raft.Leader {
			return node, nil
		}
	}

	return &models.Node{}, fmt.Errorf("no leader available")
}

func createRaftNode(id, addr string) *models.Node {
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(id)

	store, err := raftboltdb.NewBoltStore(fmt.Sprintf("%s-log.bolt", id))
	if err != nil {
		log.Fatalf("Error creating store: %v", err)
	}

	transport, err := raft.NewTCPTransport(addr, nil, 3, time.Second, nil)
	if err != nil {
		log.Fatalf("Error creating transport: %v", err)
	}

	snapshots, err := raft.NewFileSnapshotStore(".", 1, nil)
	if err != nil {
		log.Fatalf("Error creating snapshot store: %v", err)
	}

	node := &models.Node{ID: id}

	raftNode, err := raft.NewRaft(config, nil, store, store, snapshots, transport)
	if err != nil {
		log.Fatalf("Error starting Raft: %v", err)
	}

	node.Raft = raftNode
	return node
}

func (raftcluster *RaftCluster) CommitTaskToCluster(task *models.Task) error {

	node, err := raftcluster.GetLeader()
	if err != nil {
		return err
	}

	errchan := make(chan error, 1)
	fsmchan := make(chan error, 1)

	go func(node *models.Node, task *models.Task) {
		taskBytes, err := json.Marshal(task)
		if err != nil {
			errchan <- fmt.Errorf("failed to marshal the task")
			return
		}

		future := node.Raft.Apply(taskBytes, time.Second*1)
		if err := future.Error(); err != nil {
			errchan <- fmt.Errorf("failed to apply task to raft log: %v", err)
			return
		}

		response := future.Response()
		if errResp, ok := response.(error); ok {
			fsmchan <- fmt.Errorf("FSM failed to apply task: %v", errResp)
			return
		}

		errchan <- nil
		fsmchan <- nil

	}(node, task)

	if err := <-errchan; err != nil {
		return err
	}

	return <-fsmchan
}

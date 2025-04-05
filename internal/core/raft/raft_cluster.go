package raft

import (
	"encoding/json"
	"errors"
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
		addr := fmt.Sprintf("127.0.0.1:%d", 9000+i)
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

	// creating a bolt db for each node to store it's logs
	store, err := raftboltdb.NewBoltStore(fmt.Sprintf("%s-log.bolt", id))
	if err != nil {
		log.Fatalf("Error creating store: %v", err)
	}

	//creating a tcp protocol transport between nodes
	transport, err := raft.NewTCPTransport(addr, nil, 3, time.Second, nil)
	if err != nil {
		log.Fatalf("Error creating transport: %v", err)
	}

	//creating a store to store our snapshots of the latest log
	snapshots, err := raft.NewFileSnapshotStore(".", 1, nil)
	if err != nil {
		log.Fatalf("Error creating snapshot store: %v", err)
	}

	node := &models.Node{ID: id}

	//creating a new FSM
	fsm := NewFSM()

	raftNode, err := raft.NewRaft(config, fsm, store, store, snapshots, transport)
	if err != nil {
		log.Fatalf("Error starting Raft: %v", err)
	}

	node.Raft = raftNode
	return node
}

func (raftcluster *RaftCluster) CommitTaskToCluster(task *models.Task) error {

	nodechan := make(chan models.Node, 1)
	go func() error {
		duration := time.Now().Add(2 * time.Second)

		for time.Now().Before(duration) {
			leader, err := raftcluster.GetLeader()
			if err != nil {
				time.Sleep(500 * time.Millisecond) // adjustable backoff
				continue
			}
			nodechan <- *leader
			return nil
		}
		return errors.New("timeout: no leader elected")
	}()

	node := <-nodechan
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

	}(&node, task)

	if err := <-errchan; err != nil {
		return err
	}

	return <-fsmchan
}

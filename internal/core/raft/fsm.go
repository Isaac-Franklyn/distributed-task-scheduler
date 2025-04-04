package raft

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"
	"github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/ports"
	"github.com/hashicorp/raft"
)

type FSM struct {
	Db ports.DbService
}

func NewFSM() *FSM {
	return &FSM{}
}

func (fsm *FSM) Apply(logEntry *raft.Log) interface{} {

	var task models.Task

	err := json.Unmarshal(logEntry.Data, &task)
	if err != nil {
		log.Printf("FSM Apply: Failed to unmarshal task: %v", err)
		return fmt.Errorf("FSM Apply: invalid log data")
	}

	err = fsm.Db.SaveTaskToDb(&task)
	if err != nil {
		log.Printf("FSM Apply: Failed to save task in DB: %v", err)
		return fmt.Errorf("FSM Apply: database error")
	}

	log.Printf("FSM Apply: Successfully saved task %s to DB", task.ID)
	return nil
}	

func (fsm *FSM) Snapshot() (raft.FSMSnapshot, error) {
	log.Println("FSM Snapshot: Snapshot requested but not implemented")
	return &NoOpSnapshot{}, nil
}

func (fsm *FSM) Restore(snapshot io.ReadCloser) error {
	log.Println("FSM Restore: Restore requested but not implemented")
	return nil
}

type NoOpSnapshot struct{}

func (s *NoOpSnapshot) Persist(sink raft.SnapshotSink) error {
	return nil
}

// Release does nothing
func (s *NoOpSnapshot) Release() {}

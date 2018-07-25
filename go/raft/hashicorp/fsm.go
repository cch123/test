package main

import (
	"io"

	"github.com/hashicorp/raft"
)

// FSM ...
type FSM struct {
	state state
}

// NewFSM ...
func NewFSM() *FSM {
	return &FSM{state: first}
}

// Apply ...
func (f *FSM) Apply(r *raft.Log) interface{} {
	f.state.transit(state(r.Data))
	return string(f.state)
}

// Snapshot ...
func (f *FSM) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

// Restore ...
func (f *FSM) Restore(io.ReadCloser) error {
	return nil
}

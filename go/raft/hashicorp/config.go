package main

import (
	"fmt"

	"github.com/hashicorp/raft"
)

var rafts map[string]*raft.Raft

func init() {
	rafts = make(map[string]*raft.Raft)
}

// create in-memory nodes and connect them
func config(num int) {
	conf := raft.DefaultConfig()
	//snapshotStore := raft.DiscardSnapshotStore{}

	addrs := []raft.ServerAddress{}
	transports := []*raft.InmemTransport{}

	for i := 0; i < num; i++ {
		addr, transport := raft.NewInmemTransport("")
		addrs = append(addrs, addr)
		transports = append(transports, transport)
	}

	// peerStore := &raft.StableStore
	memStore := raft.NewInmemStore()
	memSnapStore := raft.NewInmemSnapshotStore()

	for i := 0; i < num; i++ {
		// connect to each other
		for j := 0; j < num; j++ {
			if i != j {
				transports[i].Connect(addrs[j], transports[j])
			}
		}

		r, err := raft.NewRaft(conf, NewFSM(),
			memStore, memStore, memSnapStore, transports[i])

		if err != nil {
			fmt.Println(err)
			panic(1)
		}

		for i := 0; i < num; i++ {
			r.AddPeer(addrs[i])
		}

		rafts[string(addrs[i])] = r
	}
}

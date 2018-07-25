package main

import (
	"fmt"
	"time"

	"github.com/hashicorp/raft"
)

var rafts map[string]*raft.Raft

func init() {
	rafts = make(map[string]*raft.Raft)
}

// create in-memory nodes and connect them
func config(num int) {
	//snapshotStore := raft.DiscardSnapshotStore{}

	addrs := []raft.ServerAddress{}
	transports := []*raft.InmemTransport{}

	for i := 0; i < num; i++ {
		addr, transport := raft.NewInmemTransport("")
		addrs = append(addrs, addr)
		transports = append(transports, transport)
	}

	store := raft.NewInmemStore()
	snaps := raft.NewInmemSnapshotStore()
	var members raft.Configuration
	for i := 0; i < num; i++ {
		conf := raft.DefaultConfig()
		conf.LocalID = raft.ServerID("test cluster" + fmt.Sprint(i))
		addr, _ := raft.NewInmemTransport("")
		members.Servers = append(members.Servers, raft.Server{
			Suffrage: raft.Voter,
			ID:       conf.LocalID,
			Address:  addr,
		})
	}

	for i := 0; i < num; i++ {
		_, trans := raft.NewInmemTransport("")
		conf := raft.DefaultConfig()
		conf.LocalID = raft.ServerID("test cluster" + fmt.Sprint(i))
		err := raft.BootstrapCluster(conf, store, store, snaps, trans, members)
		if err != nil {
			println(err)
		}

		raft, err := raft.NewRaft(conf, NewFSM(), store, store, snaps, trans)
		if err != nil {
			println(err)
		}

		timeout := time.After(10 * time.Second)
		go func() {
			for {
				if raft.Leader() != "" {
					break
				}

				select {
				case <-raft.LeaderCh():
				case <-time.After(1 * time.Second):
					// Need to poll because we might have missed the first
					// go with the leader channel.
				case <-timeout:
					println("ffdfds")
				}
			}
		}()
	}

}

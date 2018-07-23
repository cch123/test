package main

import "github.com/hashicorp/raft"

func main() {
	conf := raft.Config{}
	logs := raft.Log{}
	stable := raft.StableStore{}
	raft.BootstrapCluster(conf, logs, stable, snaps, trans, configuration)
}

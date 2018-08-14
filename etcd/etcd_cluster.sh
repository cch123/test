#!/bin/bash
etcd -name etcd01 --data-dir /home/caochunhui/data1 \
 --initial-advertise-peer-urls http://127.0.0.1:2380 \
 --listen-peer-urls http://127.0.0.1:2380 \
 --listen-client-urls http://127.0.0.1:2379,http://127.0.0.1:2379 \
 --advertise-client-urls http://127.0.0.1:2379 \
 --initial-cluster-token etcd-cluster \
 --initial-cluster etcd01=http://127.0.0.1:2380,etcd02=http://127.0.0.1:32380,etcd03=http://127.0.0.1:22380\
 --initial-cluster-state new


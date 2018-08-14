#!/bin/bash

etcd -name etcd02 --data-dir /home/caochunhui/data2 \
 --initial-advertise-peer-urls http://127.0.0.1:32380 \
 --listen-peer-urls http://127.0.0.1:32380 \
 --listen-client-urls http://127.0.0.1:12379,http://127.0.0.1:12379 \
 --advertise-client-urls http://127.0.0.1:12379 \
 --initial-cluster-token etcd-cluster \
 --initial-cluster etcd01=http://127.0.0.1:2380,etcd02=http://127.0.0.1:32380,etcd03=http://127.0.0.1:22380\
 --initial-cluster-state new


#!/usr/bin/env bash

debconf-set-selections <<< 'mysql-server mysql-server/root_password password MySuperPassword'
debconf-set-selections <<< 'mysql-server mysql-server/root_password_again password MySuperPassword'
apt-get update
apt-get install -y mysql-server

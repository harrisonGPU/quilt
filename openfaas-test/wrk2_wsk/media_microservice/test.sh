#!/bin/bash

IP=130.127.133.219
WRK_BIN=../wrk
WRK_SCRIPT="lua_files/$1.lua"
CLUSTER_ID=$2
# cluster 1 IP
ENTRY_HOST="http://$IP:30080"

QPS=400

$WRK_BIN -t 1 -c 1 -d 180 -L -U \
	 -s $WRK_SCRIPT \
	 $ENTRY_HOST -R $QPS 2>/dev/null > output_$1.log

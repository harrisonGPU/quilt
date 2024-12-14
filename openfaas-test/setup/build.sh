#!/bin/bash

USER="Harrison"

function setup_k8s {
  ### setup the kubernetes cluster
  sudo chmod -R 777 /users/Harrison/.docker
  k3sup plan machine.json \
    --user $USER \
    --servers 1 \
    --server-k3s-extra-args "--disable traefik" \
    --background > bootstrap.sh
  chmod +x bootstrap.sh
  ./bootstrap.sh
#  k3sup install --ip $SERVER_IP --user $USER
#  k3sup join --ip $AGENT_IP --server-ip $SERVER_IP --user $USER
  export KUBECONFIG=`pwd`/kubeconfig
  kubectl config use-context default
  kubectl get node -o wide
}


function setup {
  setup_k8s
  cd grafana_tempo && ./build.sh setup && cd ..
  cd open-telemetry && ./build.sh setup && cd ..
  cd ingress-nginx && ./build.sh setup && cd ..
  cd openfaas && ./build.sh setup && cd ..
  cd openwhisk && ./build.sh setup && cd ..
  cd redis_memcached && ./build.sh setup && cd ..
}

function kill_k8s {
  ALL_ENGINE_NODES=$(./helper.py load_machine_info)
  for host in $ALL_ENGINE_NODES; do
    ssh -q $USER@$host -- sudo sh /usr/local/bin/k3s-killall.sh
    ssh -q $USER@$host -- sudo sh /usr/local/bin/k3s-uninstall.sh
    ssh -q $USER@$host -- sudo sh /usr/local/bin/k3s-agent-uninstall.sh
    ssh -q $USER@$host -- npx kill-port 30080 6379 27017 11211 30081 3000 30443 30442 31112 31113 8080 6443
  done
  for entry in "$(pwd)"/*
  do
    if [[ -d $entry  ]] ; then
      rm -rf $entry/*.txt $entry/*.yaml $entry/*.yml 
    fi
  done
  rm -rf kubeconfig bootstrap.sh
}

case "$1" in
setup)
    setup
    ;;
kill)
    kill_k8s
    ;;
esac


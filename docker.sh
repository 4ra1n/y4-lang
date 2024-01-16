#!/bin/bash

case "$1" in
  "stop")
    echo "stopping all containers..."
    docker kill -f $(docker ps -a -q) >/dev/null 2>&1
    ;;
  "clean")
    echo "cleaning all containers..."
    docker rm -f $(docker ps -a -q) >/dev/null 2>&1
    ;;
  "rmi")
    echo "removing all images..."
    docker rmi -f $(docker images -q) >/dev/null 2>&1
    ;;
  *)
    echo "usage: $0 {stop|clean|rmi}"
    exit 1
    ;;
esac

echo "done"

#!/bin/bash
green() {
  "$@" | GREP_COLORS='mt=01;32' grep --color .
}

red() {
  "$@" | GREP_COLORS='mt=01;31' grep --color .
}

yellow() {
  "$@" | GREP_COLORS='mt=01;93' grep --color .
}

check_command_success() {
  CODE_TO_COMPARE_TO=$2
  RETURNED_CODE=$1
  if [ $RETURNED_CODE -ne $CODE_TO_COMPARE_TO ]; then
    if [[ $2 != "" ]]; then
      red echo "$3"
    fi
    exit 1
  fi
}

build_greedy() {
  green echo "Start building word chains resolver with greedy algorithm container"
  docker build -t greedy:build . -f ./build/package/greedy/Dockerfile
  RESULT=$?
  check_command_success $RESULT 0 "Could not build word chains resolver"

  green echo "Creating container to build word chains resolver with greedy algorithm"
  docker container create --name extract_greedy greedy:build
  RESULT=$?
  check_command_success $RESULT 0 "Could not start builder container"

  green echo "Extracting binary from builder container"
  docker container cp extract_greedy:/go/src/github.com/clnbs/wordChains/greedy.bin ./greedy.bin
  RESULT=$?
  check_command_success $RESULT 0 "Could not extract binary from builder image"

  green echo "Deleting builder container"
  docker container rm -f extract_greedy
  RESULT=$?
  check_command_success $RESULT 0 "Could not remove builder container"
}

build_bfs() {
  green echo "Start building word chains resolver with BFS algorithm container"
  docker build -t bfs:build . -f ./build/package/bfs/Dockerfile
  RESULT=$?
  check_command_success $RESULT 0 "Could not build word chains resolver"

  green echo "Creating container to build word chains resolver with bfs algorithm"
  docker container create --name extract_bfs bfs:build
  RESULT=$?
  check_command_success $RESULT 0 "Could not start builder container"

  green echo "Extracting binary from builder container"
  docker container cp extract_bfs:/go/src/github.com/clnbs/wordChains/bfs.bin ./bfs.bin
  RESULT=$?
  check_command_success $RESULT 0 "Could not extract binary from builder image"

  green echo "Deleting builder container"
  docker container rm -f extract_bfs
  RESULT=$?
  check_command_success $RESULT 0 "Could not remove builder container"
}

OPTION=$1

if [ -z "$OPTION"  ]; then
  green echo "Compiling all implementation"
  build_bfs
  build_greedy
  exit 0
elif [[ "$OPTION" == "greedy" ]]; then
  green echo "Compiling greedy implementation"
  build_greedy
elif [[ "$OPTION" == "bfs" ]]; then
  green echo "Compiling BFS implementation"
  build_bfs
fi
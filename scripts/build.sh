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

build_from_docker() {
  build_target=$1

  green echo "Start building word chains solver ${build_target} container"
  docker build -t "${build_target}":build . -f ./build/package/"${build_target}"/Dockerfile
  RESULT=$?
  check_command_success $RESULT 0 "Could not build word chains solver ${build_target} container"

  green echo "Creating container to build word chains solver ${build_target}"
  docker container create --name extract "${build_target}":build
  RESULT=$?
  check_command_success $RESULT 0 "Could not start builder container"

  green echo "Extracting binary from builder container"
  docker container cp extract:/go/src/github.com/clnbs/wordChains/"${build_target}".bin ./"${build_target}".bin
  RESULT=$?
  check_command_success $RESULT 0 "Could not extract binary from builder image"

  green echo "Deleting builder container"
  docker container rm -f extract
  RESULT=$?
  check_command_success $RESULT 0 "Could not remove builder container"
}

OPTION=$1

if [ -z "$OPTION"  ]; then
  green echo "Compiling all implementation"
  build_from_docker bfs
  build_from_docker greedy
  build_from_docker astar
  exit 0
elif [[ "$OPTION" == "greedy" ]]; then
  green echo "Compiling greedy implementation"
  build_from_docker greedy
elif [[ "$OPTION" == "bfs" ]]; then
  green echo "Compiling BFS implementation"
  build_from_docker bfs
elif [[ "$OPTION" == "astar" ]]; then
  green echo "Compiling A* implementation"
  build_from_docker astar
fi
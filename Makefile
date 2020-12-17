# note: call scripts from /scripts
.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

all: greedy bfs astar

testing: ## Start all static test for this project and create a coverage file in HTML
	bash scripts/test.sh

greedy: ## Compile greedy implementation of word chains solver
	bash scripts/build.sh greedy

bfs: ## Compile BFS implementation of word chains solver
	bash scripts/build.sh bfs

astar: ## Compile A* implementation of word chains solver
	bash scripts/build.sh astar
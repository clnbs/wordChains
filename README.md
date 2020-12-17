Code Kata19: Word Chains
===

Implementation of Code Kata19 exercice. Link to the challenge [here](http://codekata.com/kata/kata19-word-chains/).

## Table of Contents

* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
  * [Start tests](#start-tests)
  * [Start each implementation](#start-each-implementation)
* [Under the hood](#under-the-hood)
  * [General methodology](#general-methodology)
  * [Greedy algorithm](#greedy-algorithm)
    * [Greedy pros](#greedy-pros)
    * [Greedy cons](#greedy-cons)
    * [How Greedy works](#how-greedy-works)
  * [Breadth first search algorithm](#breadth-first-search-algorithm)
    * [BFS pros](#bfs-pros)
    * [BFS cons](#bfs-cons)
    * [How BFS works](#how-bfs-works)
  * [Other possible algorithm](#other-possible-algorithms)
* [TODO list](#todo-list)
* [License](#license)
* [Contact](#contact)
* [Acknowledgements](#acknowledgements)


## Getting Started

### Prerequisites

In order to compile and run the different implementation, you will need :
 - Docker installed. You can find instructions [here](https://docs.docker.com/get-docker/)
 - Git
 
### Installation
1. Clone this repository in your GOPATH : 
```bash
mkdir -p $GOPATH/src/github.com/clnbs
cd $GOPATH/src/github.com/clnbs
git clone https://github.com/clnbs/autorace/wordChains
cd wordChains
```

1. Compiling implementations
```
make all
```
This command should create a binary for each implementation 
 
## Usage

### Start tests
Tests are trigger in a Docker environment in order to be compliant with the compilation environment. Tests can be easily done using the Makefile : 
```bash
make testing
```

This command will start all static tests and create an HTML file named `cover.html` in the root directory of this project. Open it with your favorite web browser.

### Start each implementation
For each implementation, you can start a specific binary in the form `<implementation_name>.bin` at the root directory of this project. To start a binary, you can simply run e.g :
```bash
./greedy.bin <path to the word lists file> <first word> <second word>
```

e.g :
```bash
./bfs.bin assets/app/small_en.txt cat dog
```

## Under the hood
 

### General methodology
The general philosophy is to build a tree containing possible words chain. The root of the tree is the starting word and the gaol is to looking for a leaf containing the ending word. This is an example with asked chain `cat` to `dog` :
```
        cat
       /   \
      /     \
   cut      cot
    |      /   \      
    |     /     \
   gut  cog     dot
         |       |
         |       |
        dog     dog
```
Once the tree is constructed, we select each leaf containing the ending word `dog` and read the tree backward. In this example, the tree tells us there are two solutions : `cat - cot - cog - dog` and `cat - cot - dot - dog`

### Greedy algorithm
I started to implement a greedy algorithm because building the entire tree use a lot of CPU's and RAM's ressources. This algorithm is "depth first" and at each tree stage, it selects the best possible options with a scoring function. For the scoring function, I check words char per char and add a point each time chars are equals, e.g : 
 - `cat` and `dog` = 0 point
 - `ogd` and `dog` = 0 point
 - `dot` and `dog` = 2 points

#### Greedy pros 
The main advantage using a greedy algorithm is CPU and RAM usage. Because of its low footprint, it gives a result at a blasting speed and is easy to implement.
 
#### Greedy cons
The greedy algorithm do not ensure to give the best solution possible, plus, it may not give any solution if it stuck itself in an impossible branch tree. For example, in the main file starting the greedy algorithm, it finds three solutions for the word chain `oil` to `bar` but cannot find a solution for `bar` to `oil`. 

#### How Greedy works 
1. We get all words of the same length from the words list.
2. We create the tree root with the starting word.
3. We create a sub words list where each word differs with only one letter from its parent node.
4. We select the best words in the sub list with the scoring function.
5. For each remaining word in the sub word list, we create a node.
6. For each node created, we start over #3 step until we get the ending word in the sub word list.
7. If we select the ending word, we register the leaf in a slice in order to read the tree backward.
8. We select paths who give the shorter answer and return them.

### Breadth-first search algorithm
After coding the Greedy algorithm, I coded the breadth-first search (BFS) algorithm but with a trick, I build the tree while running the BFS algorithm. Building the entire tree will smash any RAM and CPU, moreover if you use the german dictionary (1 648 984 words). In the algorithm, as soon as I get a solution, I store the corresponding node tree in a slice and its depth. Then, the BFS stop looking for node with a depth higher than the solution found. The node depth is corresponding to the word chain length.  

#### BFS pros
BFS is complete solution, if there is a solution, it will found it.

#### BFS cons
BFS consume more RAM and CPU ressources than some other algorithm. Due to this, it is way slower than the greedy algorithm. For some complicated word chains, it may looks for a soltuion  

#### How BFS works
1. We get all words of the same length from the words list.
2. We create the tree root with the starting word.
3. We mark the root tree node as marked.
4. We add the root tree node in, a queue. This queue stores nodes who do not have neighbors created.
5. While the queue is not empty :
6. - We pop a node from the queue
7. - If the depth of the node popped if bigger than an already registered solution, we end the process
8. - If the node popped contains the ending word, we register it as a solution and mark its depth
9. - If non of the condition above is observe, we create a sub words list where each word differs with only one letter from the popped node.
10. - We exclude from the possible words list the word already registered in previous nodes
11. - For every word remaining in the possible word list : 
12.   - We create a node containing the possible next word
13.   - If the node is not already discover, we mark it as discovered and add it to the queue
14. Start again at the step #6

### A*
After seeing BFS algorithm struggling to compute word chain, I decided to implement the A* algorithm, the best path finding algorithm out there yet. It has a super low complexity (O(b^d)) and is complete. If there is a solution, it will find it at a super speed. 

#### A* pros
A* is blasting fast and ensure to find a solution if there is any. 

#### A* cons
A* can return only one solution. A work around is possible, but I am running out of time to implement it.

#### How A* works
1. We get all words of the same length from the words list. --> it starts to look familiar ...
2. We create the tree root with the starting word.
3. We mark the node as an open set (to get computed)
4. We compute its depth and store it in a map noted as G score (map[node as key] = score as value)
5. We compute its difference from the ending word, the more they are different, the bigger the score is. We store it in a map noted as F score (map[node as key] = score as value)
6. While the open set list is not empty :
7. - We select the best current node in the open set list.
8. - If the current best is the goal, we return the solution and stop the execution.
9. - We create all neighbors of the current best 
10. - For every created neighbor of the current best :
11.   - We calculate the G score and store it in the corresponding map
12.   - We calculate the F score and store it in the corresponding map
13.   - We add the node in the open set list
14. Start again at the step #7


### Other possible algorithms
Even if the best path finding algorithm is A*, other algorithms could be used to make word chains. They all got pros and cons too, here is some example :    
 - Dijkstra : Complete but slow, A little like BFS in that case 
 - DFS : Similar to BFS but it looks for a solution in depth first. It is completely irrelevant in our case. 

## TODO list
 - Refactoring test code

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Colin Bois - <colin.bois@rocketmail.com>

Project Link: [https://github.com/clnbs/wordChains](https://github.com/clnbs/wordChains)

## Acknowledgements

* [Code Kata](http://codekata.com/)
* [Cycloid.io](https://www.cycloid.io/)
* [mieliestronk.com - word list](http://www.mieliestronk.com/corncob_lowercase.txt)
* [Wikipedia - BFS](https://en.wikipedia.org/wiki/Breadth-first_search)
* [Golang testify from stretchr](https://github.com/stretchr/testify)
* [lorenbrichter - github.com - word list](https://github.com/lorenbrichter/Words)
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
1. Getting all words of the same length from the words list.
2. Creating the tree root with the starting word.
3. Creating a sub words list where each word differs with only one letter from its parent node.
4. Selecting the best words in the sub list with the scoring function.
5. For each remaining word in the sub word list, we create a node.
6. For each node created, we start over #3 step until we get the ending word in the sub word list.
7. If we select the ending word, we register the leaf in a slice in order to read the tree backward.
8. We select paths who give the shorter answer and return them.

### Breadth-first search algorithm
[WIP]

#### BFS pros
[WIP]

#### BFS cons
[WIP]

#### How BFS works
[WIP]

### Other possible algorithms
[WIP]
 - Dijkstra
 - DFS
 - A*

## TODO list
 - Make tests and builds automated with a Makefile and Docker containers
 - This README.md file in BFS and other algorithms section

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
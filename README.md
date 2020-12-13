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
With command should create a binary for each implementation 
 
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
./greedy.bin
```

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Colin Bois - <colin.bois@rocketmail.com>

Project Link: [https://github.com/clnbs/karateChop](https://github.com/clnbs/karateChop)

## Acknowledgements

* [Code Kata](http://codekata.com/)
* [Cycloid.io](https://www.cycloid.io/)
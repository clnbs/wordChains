package wordchainsresolver

// BFSWordTreeNode struct represents words tidy in a tree
type BFSWordTreeNode struct {
	Word            string
	PreviousElement *BFSWordTreeNode
	NextElements    []*BFSWordTreeNode
}

// NewBFSWordTreeNode is BFSWordTreeNode constructor
func NewBFSWordTreeNode(word string, previous *BFSWordTreeNode) *BFSWordTreeNode {
	newNode := &BFSWordTreeNode{
		Word:            word,
		PreviousElement: previous,
		NextElements:    nil,
	}
	if previous != nil {
		previous.NextElements = append(previous.NextElements, newNode)
	}
	return newNode
}

//Depth return a node's depth in the tree
func (node *BFSWordTreeNode) Depth() int {
	depth := 1
	tmpNode := node
	for tmpNode.PreviousElement != nil {
		depth++
		tmpNode = tmpNode.PreviousElement
	}
	return depth
}

// GetSolution return the word chain from the current node
// by looking at its parent node until it reach the root node
func (node *BFSWordTreeNode) GetSolution() []string {
	var wordChains []string
	tmpNode := node
	wordChains = append(wordChains, tmpNode.Word)
	for tmpNode.PreviousElement != nil {
		tmpNode = tmpNode.PreviousElement
		wordChains = append(wordChains, tmpNode.Word)
	}

	return flipStringSlice(wordChains)
}

// BFSQueue is a BFSWordTreeNode FIFO queue
type BFSQueue struct {
	words []*BFSWordTreeNode
}

// Pop return the first element of the queue
func (bfsQ *BFSQueue) Pop() *BFSWordTreeNode {
	if len(bfsQ.words) == 0 {
		return nil
	}
	var element *BFSWordTreeNode
	element, bfsQ.words = bfsQ.words[0], bfsQ.words[1:]
	return element
}

// Add adds a element at the back of the queue
func (bfsQ *BFSQueue) Add(element *BFSWordTreeNode) {
	bfsQ.words = append(bfsQ.words, element)
}

// Len return queue's length
func (bfsQ *BFSQueue) Len() int {
	return len(bfsQ.words)
}

// BFSSolver is a implementation of Solver interface in order to find
//// word chains with a BFS algorithm
type BFSSolver struct {
	wordsList         []string
	usefulWords       []string
	from              string
	to                string
	queue             *BFSQueue
	wordTree          *BFSWordTreeNode
	solutions         []*BFSWordTreeNode
	bestSolutionDepth int
	discovered        map[*BFSWordTreeNode]interface{}
}

// NewBFSSolver is a simple BFSSolver constructor
func NewBFSSolver() *BFSSolver {
	return &BFSSolver{
		bestSolutionDepth: int(^uint(0) >> 1),
		discovered:        make(map[*BFSWordTreeNode]interface{}),
		queue:             &BFSQueue{},
	}
}

// NewBFSSolverWithParams is also a BSFSolver constructor but with params
// input : the first word of the futur word chains, the ending word of the futur word chains,
// the word list database
// /!\ Warning, using this constructor is unsafe and should be used in a testing purpose
func NewBFSSolverWithParams(from, to string, wordList []string) *BFSSolver {
	return &BFSSolver{
		wordsList:         wordList,
		from:              from,
		to:                to,
		queue:             &BFSQueue{},
		wordTree:          nil,
		solutions:         nil,
		bestSolutionDepth: int(^uint(0) >> 1),
		discovered:        make(map[*BFSWordTreeNode]interface{}),
	}
}

// FindWordChains implements the Solver interface. The BFS solver generate word chains
// by looking for the best solutions in a tree, breadth first. It is a complete algorithm :
// if there is a solution, BFS will find it
func (bfs *BFSSolver) FindWordChains(from string, to string, wordList []string) ([][]string, error) {
	if len(from) != len(to) {
		return nil, ErrorWordLengthDoesNotMatch
	}
	bfs.from = from
	bfs.to = to
	bfs.wordsList = wordList

	bfs.getUsefulWordOnly()
	bfs.wordTree = NewBFSWordTreeNode(from, nil)
	bfs.solveBFS()

	var solutions [][]string
	for _, node := range bfs.solutions {
		solutions = append(solutions, node.GetSolution())
	}
	solutions = getBestSolution(solutions)
	bfs.Clean()
	return solutions, nil
}

func (bfs *BFSSolver) solveBFS() {
	// mark tree's root as discovered
	bfs.discovered[bfs.wordTree] = nil
	bfs.queue.Add(bfs.wordTree)

	for bfs.queue.Len() != 0 {
		node := bfs.queue.Pop()
		if node.Word == bfs.to {
			nodeDepth := node.Depth()
			if nodeDepth <= bfs.bestSolutionDepth {
				bfs.bestSolutionDepth = nodeDepth
				bfs.solutions = append(bfs.solutions, node)
			}
			if nodeDepth > bfs.bestSolutionDepth {
				return
			}
		}

		possibleWords := bfs.listPossibleNextWords(node.Word)
		alreadyRegisteredWords := node.GetSolution()
		possibleWords = excludeStringsFromStrings(possibleWords, alreadyRegisteredWords)
		for _, word := range possibleWords {
			newNode := NewBFSWordTreeNode(word, node)
			if _, ok := bfs.discovered[newNode]; !ok {
				bfs.discovered[newNode] = nil
				bfs.queue.Add(newNode)
			}
		}
	}
}

func (bfs *BFSSolver) getUsefulWordOnly() {
	wordLength := len(bfs.from)
	for _, word := range bfs.wordsList {
		if len(word) == wordLength && word != bfs.from {
			bfs.usefulWords = append(bfs.usefulWords, word)
		}
	}
}

func (bfs *BFSSolver) listPossibleNextWords(word string) []string {
	var possibleNewWords []string
	for _, nextWord := range bfs.usefulWords {
		if isPossibleNextWord(nextWord, word) {
			possibleNewWords = append(possibleNewWords, nextWord)
		}
	}
	return possibleNewWords
}

// Clean delete all data stored in the current BFSSolver instance
func (bfs *BFSSolver) Clean() {
	bfs.wordsList = nil
	bfs.usefulWords = nil
	bfs.from = ""
	bfs.to = ""
	bfs.queue = &BFSQueue{}
	bfs.wordTree = nil
	bfs.solutions = []*BFSWordTreeNode{}
	bfs.bestSolutionDepth = int(^uint(0) >> 1)
	bfs.discovered = make(map[*BFSWordTreeNode]interface{})
}

package wordchainsresolver

// AStarNode struct represents words tidy in a tree node
type AStarNode struct {
	word     string
	previous *AStarNode
}

// NewAStarNode is the AStarNode constructor
func NewAStarNode(word string, previous *AStarNode) *AStarNode {
	return &AStarNode{
		word:     word,
		previous: previous,
	}
}

//Depth return a node's depth in the tree
func (node *AStarNode) Depth() int {
	score := 0
	tmpNode := node
	for tmpNode != nil {
		score++
		tmpNode = tmpNode.previous
	}
	return score
}

// GetSolution return the word chain from the current node
// by looking at its parent node until it reach the root node
func (node *AStarNode) GetSolution() []string {
	var wordChains []string
	tmpNode := node
	for tmpNode != nil {
		wordChains = append(wordChains, tmpNode.word)
		tmpNode = tmpNode.previous
	}
	return flipStringSlice(wordChains)
}

// AStarSolver is a implementation of Solver interface in order to find
// word chains with a A* algorithm
type AStarSolver struct {
	tree        *AStarNode
	nodeGScore  map[*AStarNode]int
	nodeFScore  map[*AStarNode]int
	openSet     map[*AStarNode]interface{}
	wordList    []string
	usefulWords []string
	from        string
	to          string
}

// NewAStarSolver is a simple AStarSolver constructor
func NewAStarSolver() *AStarSolver {
	return &AStarSolver{
		tree:        nil,
		nodeGScore:  make(map[*AStarNode]int),
		nodeFScore:  make(map[*AStarNode]int),
		openSet:     make(map[*AStarNode]interface{}),
		wordList:    nil,
		usefulWords: nil,
	}
}

// FindWordChains implements the Solver interface. The A* solver generate word chains
// by looking for the best solutions in a tree. It is a complete algorithm :
// if there is a solution, A* will find it
func (a *AStarSolver) FindWordChains(from string, to string, wordList []string) ([][]string, error) {
	if len(from) != len(to) {
		return nil, ErrorWordLengthDoesNotMatch
	}
	defer a.Clean()
	// A* initialisation, go see README.md for more information
	goal := NewAStarNode(to, nil)
	head := NewAStarNode(from, nil)
	a.from = from
	a.to = to
	a.wordList = wordList
	a.getUsefulWordsOnly()

	a.openSet[head] = nil
	a.nodeGScore[head] = head.Depth()
	a.nodeFScore[head] = a.getScoreFromGoal(head)

	// A* main loop
	for len(a.openSet) != 0 {
		// can be improved with priorityQueue
		current := a.getCurrentBestNode()
		if current.word == goal.word {
			return [][]string{current.GetSolution()}, nil
		}
		delete(a.openSet, current)

		neighbors := a.createNeighbors(current)
		for _, neighbor := range neighbors {
			score := neighbor.Depth()
			a.nodeGScore[neighbor] = score
			fScore := a.getScoreFromGoal(neighbor) + neighbor.Depth()
			a.nodeFScore[neighbor] = fScore
			if _, ok := a.openSet[neighbor]; !ok {
				a.openSet[neighbor] = nil
			}
		}
	}
	return nil, nil
}

func (a *AStarSolver) getCurrentBestNode() *AStarNode {
	currentBest := &AStarNode{}
	currentBestScore := int(^uint(0) >> 1)
	for node, score := range a.nodeFScore {
		_, isSelectable := a.openSet[node]
		if score < currentBestScore && isSelectable {
			currentBest = node
			currentBestScore = score
		}
	}
	return currentBest
}

func (a *AStarSolver) createNeighbors(node *AStarNode) []*AStarNode {
	var neighbor []*AStarNode
	for _, nextWord := range a.usefulWords {
		if isPossibleNextWord(nextWord, node.word) {
			neighbor = append(neighbor, NewAStarNode(nextWord, node))
		}
	}
	return neighbor
}

func (a *AStarSolver) getUsefulWordsOnly() {
	wordLength := len(a.from)
	for _, word := range a.wordList {
		if len(word) == wordLength && word != a.from {
			a.usefulWords = append(a.usefulWords, word)
		}
	}
}

func (a *AStarSolver) getScoreFromGoal(node *AStarNode) int {
	return len(a.to) - getScoreBetweenTwoWord(node.word, a.to)
}

// Clean delete all data stored in the current AStarSolver instance
func (a *AStarSolver) Clean() {
	a.tree = nil
	a.nodeGScore = make(map[*AStarNode]int)
	a.nodeFScore = make(map[*AStarNode]int)
	a.openSet = make(map[*AStarNode]interface{})
	a.wordList = []string{}
	a.usefulWords = []string{}
}

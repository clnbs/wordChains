package wordchainsresolver

// GreedyWordTreeNode struct represents words tidy in a tree
type GreedyWordTreeNode struct {
	Word            string
	ScoreToGoal     int
	PreviousElement *GreedyWordTreeNode
	NextElements    []*GreedyWordTreeNode
}

// NewGreedyWordTreeElement is GreedyWordTreeNode constructor
// input : a word to be push in the tree, its score from the final word, its previous element
func NewGreedyWordTreeElement(word string, score int, previous *GreedyWordTreeNode) *GreedyWordTreeNode {
	return &GreedyWordTreeNode{
		Word:            word,
		ScoreToGoal:     score,
		PreviousElement: previous,
		NextElements:    nil,
	}
}

func (node *GreedyWordTreeNode) extractSolutionFromNode() []string {
	var wordChains []string
	tmpNode := node
	wordChains = append(wordChains, tmpNode.Word)
	for tmpNode.PreviousElement != nil {
		tmpNode = tmpNode.PreviousElement
		wordChains = append(wordChains, tmpNode.Word)
	}
	return flipStringSlice(wordChains)
}

func (node *GreedyWordTreeNode) getNodeDepth() int {
	depth := 1
	tmpNode := node
	for tmpNode.PreviousElement != nil {
		depth++
		tmpNode = tmpNode.PreviousElement
	}
	return depth
}

// GreedySolver is a implementation of Solver interface in order to find
// word chains with a greedy algorithm
type GreedySolver struct {
	wordList             []string
	from                 string
	to                   string
	usefulWords          []string
	wordTree             *GreedyWordTreeNode
	matchingWordNode     []*GreedyWordTreeNode
	solutionFoundAtDepth int
	maxDepth             int
}

// NewGreedySolver is a simple GreedySolver constructor
func NewGreedySolver() *GreedySolver {
	return &GreedySolver{
		solutionFoundAtDepth: int(^uint(0) >> 1),
	}
}

// NewGreedySolverWithParams is a GreedySolver constructor too but with params
// input : the first word of the futur word chains, the ending word of the futur word chains,
// the word list database
// /!\ Warning, using this constructor is unsafe and should be used in a testing purpose
func NewGreedySolverWithParams(from string, to string, wordList []string) *GreedySolver {
	return &GreedySolver{
		wordList:             wordList,
		from:                 from,
		to:                   to,
		usefulWords:          nil,
		wordTree:             nil,
		matchingWordNode:     nil,
		solutionFoundAtDepth: int(^uint(0) >> 1),
		maxDepth:             len(from) * 3,
	}
}

// FindWordChains implements the Solver interface. The greedy solver generate a word chain
// using the greedy algorithm. It is not complete so it may not give any result
func (greedy *GreedySolver) FindWordChains(from string, to string, wordList []string) ([][]string, error) {
	if len([]rune(from)) != len([]rune(to)) {
		return nil, ErrorWordLengthDoesNotMatch
	}
	greedy.from = from
	greedy.to = to
	greedy.wordList = wordList
	greedy.maxDepth = len(from) * 3

	greedy.getUsefulWordOnly()

	solutions := greedy.getPath()
	greedy.Clean()
	return getBestSolution(solutions), nil
}

func (greedy *GreedySolver) getPath() [][]string {
	head := NewGreedyWordTreeElement(greedy.from, getScoreBetweenTwoWord(greedy.from, greedy.to), nil)

	var wordChainsList [][]string

	greedy.wordTree = greedy.generateTree(head, []string{greedy.from})

	for _, solutionNode := range greedy.matchingWordNode {
		wordChainSolution := solutionNode.extractSolutionFromNode()
		wordChainsList = append(wordChainsList, wordChainSolution)
	}
	return wordChainsList
}

func (greedy *GreedySolver) generateTree(head *GreedyWordTreeNode, wordList []string) *GreedyWordTreeNode {
	// Ending condition
	if head.Word == greedy.to {
		greedy.solutionFoundAtDepth = head.getNodeDepth()
		greedy.matchingWordNode = append(greedy.matchingWordNode, head)
		return head
	}
	if head.getNodeDepth() > greedy.solutionFoundAtDepth {
		return head
	}

	possibleNextWords := greedy.listPossibleNextWords(head.Word)
	possibleNextWords = excludeStringsFromStrings(possibleNextWords, wordList)
	numberOfChildAdded := 1
	targetedScore := head.ScoreToGoal + 1
	head, numberOfChildAdded = greedy.createPopulation(head, possibleNextWords, wordList, targetedScore)
	depth := head.getNodeDepth()
	if numberOfChildAdded == 0 && depth < greedy.maxDepth {
		head, numberOfChildAdded = greedy.createPopulation(head, possibleNextWords, wordList, targetedScore-1)
	}

	return head
}

func (greedy *GreedySolver) createPopulation(head *GreedyWordTreeNode, possibleNextWords, wordList []string, targetedScore int) (*GreedyWordTreeNode, int) {
	numberOfNodeCreated := 0
	for _, word := range possibleNextWords {
		scoreFromGoal := getScoreBetweenTwoWord(word, greedy.to)
		if scoreFromGoal == targetedScore {
			numberOfNodeCreated++
			newNode := NewGreedyWordTreeElement(word, scoreFromGoal, head)
			wordList = append(wordList, word)
			newNode = greedy.generateTree(newNode, wordList)
			head.NextElements = append(head.NextElements, newNode)
		}
	}
	return head, numberOfNodeCreated
}

func (greedy *GreedySolver) getUsefulWordOnly() {
	wordLength := len(greedy.from)
	for _, word := range greedy.wordList {
		if len(word) == wordLength && word != greedy.from {
			greedy.usefulWords = append(greedy.usefulWords, word)
		}
	}
}

func (greedy *GreedySolver) listPossibleNextWords(word string) []string {
	var possibleNewWords []string
	for _, nextWord := range greedy.usefulWords {
		if getScoreBetweenTwoWord(nextWord, word) == len(word)-1 {
			possibleNewWords = append(possibleNewWords, nextWord)
		}
	}
	return possibleNewWords
}

// Clean delete all data stored in the current GreedySolver instance
func (greedy *GreedySolver) Clean() {
	greedy.from = ""
	greedy.to = ""
	greedy.usefulWords = nil
	greedy.wordTree = nil
	greedy.matchingWordNode = nil
	greedy.solutionFoundAtDepth = int(^uint(0) >> 1)
}

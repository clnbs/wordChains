package wordchainsresolver

type BFSWordTreeNode struct {
	Word            string
	PreviousElement *BFSWordTreeNode
	NextElements    []*BFSWordTreeNode
}

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

func (node *BFSWordTreeNode) Depth() int {
	depth := 1
	tmpNode := node
	for tmpNode.PreviousElement != nil {
		depth++
		tmpNode = tmpNode.PreviousElement
	}
	return depth
}

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

type BFSQueue struct {
	words []*BFSWordTreeNode
}

func (bfsQ *BFSQueue) Pop() *BFSWordTreeNode {
	if len(bfsQ.words) == 0 {
		return nil
	}
	var element *BFSWordTreeNode
	element, bfsQ.words = bfsQ.words[0], bfsQ.words[1:]
	return element
}

func (bfsQ *BFSQueue) Add(element *BFSWordTreeNode) {
	bfsQ.words = append(bfsQ.words, element)
}

func (bfsQ *BFSQueue) Len() int {
	return len(bfsQ.words)
}

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

func NewBFSSolver() *BFSSolver {
	return &BFSSolver{
		bestSolutionDepth: int(^uint(0) >> 1),
		discovered:        make(map[*BFSWordTreeNode]interface{}),
		queue:             &BFSQueue{},
	}
}

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

package wordchainsresolver

import (
	"errors"
)

var (
	// ErrorWordLengthDoesNotMatch is trigger when the words enter to create a word chain
	// are not the same size
	ErrorWordLengthDoesNotMatch = errors.New("greedy solver : word length does not match")
)

// GreedyWordTreeNode struct represents words tidy in a tree
type GreedyWordTreeNode struct {
	Word            string
	ScoreToGoal     int
	PreviousElement *GreedyWordTreeNode
	NextElements    []*GreedyWordTreeNode
}

// NewWordTreeElement is a GreedyWordTreeNode constructor
func NewWordTreeElement(word string, score int, previous *GreedyWordTreeNode) *GreedyWordTreeNode {
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

// GreedySolver is a implementation of Solver interface
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
	if len(from) != len(to) {
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
	head := NewWordTreeElement(greedy.from, computeScoreBetweenTwoWord(greedy.from, greedy.to), nil)

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
		scoreFromGoal := computeScoreBetweenTwoWord(word, greedy.to)
		if scoreFromGoal == targetedScore {
			numberOfNodeCreated++
			newNode := NewWordTreeElement(word, scoreFromGoal, head)
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
		if computeScoreBetweenTwoWord(nextWord, word) == len(word)-1 {
			possibleNewWords = append(possibleNewWords, nextWord)
		}
	}
	return possibleNewWords
}

// Clean delete all data from previous run
func (greedy *GreedySolver) Clean() {
	greedy.from = ""
	greedy.to = ""
	greedy.usefulWords = nil
	greedy.wordTree = nil
	greedy.matchingWordNode = nil
	greedy.solutionFoundAtDepth = int(^uint(0) >> 1)
}

func computeScoreBetweenTwoWord(word1, word2 string) int {
	if len(word1) != len(word2) {
		return 0
	}
	var score int
	word1Chars := []rune(word1)
	word2Chars := []rune(word2)
	for index, char := range word1Chars {
		if char == word2Chars[index] {
			score++
		}
	}
	return score
}

func excludeStringsFromStrings(strs, bannedWords []string) []string {
	var strsWithoutBannedWords []string
	for _, str := range strs {
		if !isWordInList(str, bannedWords) {
			strsWithoutBannedWords = append(strsWithoutBannedWords, str)
		}
	}
	return strsWithoutBannedWords
}

func isWordInList(word string, wordList []string) bool {
	for _, wordInList := range wordList {
		if word == wordInList {
			return true
		}
	}
	return false
}

func flipStringSlice(strSlice []string) []string {
	flipStrSlice := make([]string, len(strSlice))
	for index := range strSlice {
		flipStrSlice[len(strSlice)-1-index] = strSlice[index]
	}
	return flipStrSlice
}

func getBestSolution(solutions [][]string) [][]string {
	bestScore := int(^uint(0) >> 1)
	var bestSolutions [][]string
	for _, solution := range solutions {
		if len(solution) < bestScore {
			bestScore = len(solution)
		}
	}
	for _, solution := range solutions {
		if len(solution) == bestScore {
			bestSolutions = append(bestSolutions, solution)
		}
	}
	return bestSolutions
}

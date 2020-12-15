package wordchainsresolver

import "errors"

var (
	// ErrorWordLengthDoesNotMatch is trigger when the words enter to create a word chain
	// are not the same size
	ErrorWordLengthDoesNotMatch = errors.New("solver : word length does not match")
)

// Factory handle everything linked to loading data
type Factory interface {
	LoadDB() ([]string, error)
}

// Solver handle calculus part of the word chains problem
type Solver interface {
	FindWordChains(string, string, []string) ([][]string, error)
}

// WordChainsResolver wrap Solver and Factory interfaces by holding
// the word lis to process
type WordChainsResolver struct {
	solver   Solver
	factory  Factory
	wordList []string
}

// NewWordChainsResolver WordChainsResolver struct constructor
func NewWordChainsResolver(solver Solver, factory Factory) *WordChainsResolver {
	return &WordChainsResolver{solver: solver, factory: factory}
}

// LoadDB Factory wrapper
func (wcr *WordChainsResolver) LoadDB() error {
	var err error
	wcr.wordList, err = wcr.factory.LoadDB()
	if err != nil {
		return err
	}
	return nil
}

// Solve Solver wrapper
func (wcr *WordChainsResolver) Solve(from, to string) ([][]string, error) {
	return wcr.solver.FindWordChains(from, to, wcr.wordList)
}

// Helpers

func getScoreBetweenTwoWord(word1, word2 string) int {
	if len(word1) != len(word2) || len(word1) == 0 {
		return -1
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

func isPossibleNextWord(word1, word2 string) bool {
	if len(word1) == 0 {
		return false
	}
	score := getScoreBetweenTwoWord(word1, word2)
	return score == len(word1)-1
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

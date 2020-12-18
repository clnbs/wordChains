package wordchainsresolver

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockFactory struct {
}

func (factory *MockFactory) LoadDB() ([]string, error) {
	return []string{"cat", "cot", "cog", "dog"}, nil
}

type MockBadFactory struct {
}

func (badFactory *MockBadFactory) LoadDB() ([]string, error) {
	return nil, errors.New("from mock : I am a bad factory, mu-ah-ah-ah-aha ")
}

type MockSolver struct {
}

func (solver *MockSolver) FindWordChains(string, string, []string) ([][]string, error) {
	return [][]string{{"cat", "cot", "cog", "dog"}, {"cat", "cot", "dot", "dog"}}, nil
}

func GeneralWordChainsResolverTest(solver Solver, factory Factory, t *testing.T) {
	expectedResult := [][]string{{"cat", "cot", "cog", "dog"}, {"cat", "cot", "dot", "dog"}}
	wcr := NewWordChainsResolver(solver, factory)
	err := wcr.LoadDB()
	assert.Nil(t, err)
	result, err := wcr.Solve("cat", "dog")
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	_, err = wcr.Solve("www", "dog")
	assert.NotNil(t, err)
}

func TestNewWordChainsResolver(t *testing.T) {
	solver := &MockSolver{}
	factory := &MockFactory{}
	wcr := NewWordChainsResolver(solver, factory)
	assert.Equal(t, wcr.solver, solver)
	assert.Equal(t, wcr.factory, factory)
}

func TestWordChainsResolver_LoadDB(t *testing.T) {
	expectedResult := []string{"cat", "cot", "cog", "dog"}
	wcr := NewWordChainsResolver(&MockSolver{}, &MockFactory{})
	err := wcr.LoadDB()
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, wcr.wordList)

	wcr = NewWordChainsResolver(&MockSolver{}, &MockBadFactory{})
	err = wcr.LoadDB()
	assert.NotNil(t, err)
}

func TestWordChainsResolver_IsWordInDB(t *testing.T) {
	wcr := NewWordChainsResolver(&MockSolver{}, &MockFactory{})
	err := wcr.LoadDB()
	assert.Nil(t, err)

	assert.Equal(t, true, wcr.IsWordInDB("cat"))
	assert.Equal(t, false, wcr.IsWordInDB("www"))
}

func TestWordChainsResolver_Solve(t *testing.T) {
	GeneralWordChainsResolverTest(&MockSolver{}, &MockFactory{}, t)
}

func TestExtractSolutionFromNode(t *testing.T) {
	head := NewGreedyWordTreeElement("test", 0, nil)
	node := NewGreedyWordTreeElement("test_depth_2", 0, head)
	expected := []string{"test", "test_depth_2"}
	assert.Equal(t, expected, node.extractSolutionFromNode())
}

func TestIsWordInList(t *testing.T) {
	wordsList := []string{"one", "two", "three"}
	assert.Equal(t, true, isWordInList("one", wordsList))
	assert.Equal(t, false, isWordInList("four", wordsList))
}

func TestExcludeStringsFromStrings(t *testing.T) {
	wordsList := []string{"one", "two", "three", "four"}
	bannedWords := []string{"two", "four"}
	expected := []string{"one", "three"}
	result := excludeStringsFromStrings(wordsList, bannedWords)
	assert.Equal(t, expected, result)
}

type GetScoreTestCase struct {
	word1    string
	word2    string
	expected int
}

func TestGetScoreBetweenTwoWord(t *testing.T) {
	testCases := []GetScoreTestCase{
		{
			word1:    "abc",
			word2:    "dbz",
			expected: 1,
		},
		{
			word1:    "abc",
			word2:    "hxh",
			expected: 0,
		},
		{
			word1:    "abc",
			word2:    "abc",
			expected: 3,
		},
		{
			word1:    "dbz",
			word2:    "hxh",
			expected: 0,
		},
		{
			word1:    "abc",
			word2:    "abcd",
			expected: -1,
		},
		{
			word1:    "",
			word2:    "",
			expected: -1,
		},
	}
	for _, test := range testCases {
		assert.Equal(t, test.expected, getScoreBetweenTwoWord(test.word1, test.word2))
	}
}

type IsPossibleNextWordTestCase struct {
	word1    string
	word2    string
	expected bool
}

func TestIsPossibleNextWord(t *testing.T) {
	testCases := []IsPossibleNextWordTestCase{
		{
			word1:    "abc",
			word2:    "abd",
			expected: true,
		},
		{
			word1:    "abc",
			word2:    "acb",
			expected: false,
		},
		{
			word1:    "abc",
			word2:    "abcd",
			expected: false,
		},
		{
			word1:    "acb",
			word2:    "acd",
			expected: true,
		},
		{
			word1:    "",
			word2:    "",
			expected: false,
		},
	}
	for _, test := range testCases {
		assert.Equal(t, test.expected, isPossibleNextWord(test.word1, test.word2), "checking words "+test.word1+" and "+test.word2)
	}
}

func TestGetBestSolution(t *testing.T) {
	expected := [][]string{
		{"1", "2", "3"},
		{"a", "b", "c"},
	}
	testingValues := [][]string{
		{"1", "2", "3"},
		{"a", "b", "c"},
		{"a", "b", "1", "2"},
	}
	result := getBestSolution(testingValues)
	assert.Equal(t, expected, result)
}

func TestFlipStringSlice(t *testing.T) {
	expected := []string{"1", "2", "3"}
	toFormat := []string{"3", "2", "1"}
	result := flipStringSlice(toFormat)
	assert.Equal(t, expected, result)
}

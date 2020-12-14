package wordchainsresolver

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockWordsList_TestListPossibleNextWords = []string{
		"cat",
		"cot",
		"cut",
	}
	mockWordsList_TestGetUsefulWordOnly = []string{
		"cat",
		"cot",
		"cut",
		"dummy",
	}
)

func buildMockWordsTree() *WordTreeNode {
	head := NewWordTreeElement("cat", 0, nil)
	head.NextElements = append(head.NextElements, NewWordTreeElement("cot", 1, head))
	head.NextElements[0].NextElements = append(head.NextElements[0].NextElements,
		NewWordTreeElement("cog", 2, head.NextElements[0]),
		NewWordTreeElement("dot", 2, head.NextElements[0]),
	)
	head.NextElements[0].NextElements[0].NextElements = append(head.NextElements[0].NextElements[0].NextElements,
		NewWordTreeElement("dog", 3, head.NextElements[0].NextElements[0]))
	head.NextElements[0].NextElements[1].NextElements = append(head.NextElements[0].NextElements[1].NextElements,
		NewWordTreeElement("dog", 3, head.NextElements[0].NextElements[1]))
	return head
}

func TestGreedySolver_FindWordChains(t *testing.T) {
	solver := NewGreedySolver()
	factory := NewFileLoaderFactory(os.Getenv("GOPATH") + "/src/github.com/clnbs/wordChains/assets/app/wordlist.txt")
	GeneralWordChainsResolverTest(solver, factory, t)
	solver = NewGreedySolver()
	_, err := solver.FindWordChains("dummy", "to", []string{})
	assert.NotNil(t, err)
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

func TestGetNodeDepth(t *testing.T) {
	head := NewWordTreeElement("test", 0, nil)
	leaf := NewWordTreeElement("test_depth_2", 0, head)
	expectedHeadDepth := 1
	expectedLeafDepth := 2
	assert.Equal(t, expectedHeadDepth, getNodeDepth(head))
	assert.Equal(t, expectedLeafDepth, getNodeDepth(leaf))
}

func TestExtractSolutionFromLeaf(t *testing.T) {
	head := NewWordTreeElement("test", 0, nil)
	leaf := NewWordTreeElement("test_depth_2", 0, head)
	expected := []string{"test", "test_depth_2"}
	assert.Equal(t, expected, extractSolutionFromLeaf(leaf))
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

func TestComputeScoreBetweenTwoWord(t *testing.T) {
	wordsList := []string{"abc", "dbz", "hxh", "abc", "abcd"}
	expectedResult := []int{1, 0, 3, 0, 0}
	result := make([]int, 5)
	result[0] = computeScoreBetweenTwoWord(wordsList[0], wordsList[1])
	result[1] = computeScoreBetweenTwoWord(wordsList[0], wordsList[2])
	result[2] = computeScoreBetweenTwoWord(wordsList[0], wordsList[3])
	result[3] = computeScoreBetweenTwoWord(wordsList[1], wordsList[2])
	result[4] = computeScoreBetweenTwoWord(wordsList[0], wordsList[4])
	assert.Equal(t, expectedResult, result)
}

func TestListPossibleNextWords(t *testing.T) {
	solver := NewGreedySolverWithParams("too", "too", mockWordsList_TestListPossibleNextWords)
	solver.getUsefulWordOnly()
	expected := []string{"cot", "cut"}
	result := solver.listPossibleNextWords("cat")
	assert.Equal(t, expected, result)
}

func TestGetUsefulWordOnly(t *testing.T) {
	solver := NewGreedySolverWithParams("too", "too", mockWordsList_TestGetUsefulWordOnly)
	expected := []string{
		"cat",
		"cot",
		"cut",
	}
	solver.getUsefulWordOnly()
	assert.Equal(t, expected, solver.usefulWords)
}

func TestGenerateTree(t *testing.T) {
	solver := NewGreedySolverWithParams("cat", "dog", []string{"cat", "cot", "cog", "dog", "dot"})
	solver.getUsefulWordOnly()
	head := NewWordTreeElement("cat", 0, nil)
	result := solver.generateTree(head, []string{"cat"})
	expected := buildMockWordsTree()
	assert.Equal(t, expected, result)
}

func ExampleGreedySolver_FindWordChains() {
	wordsList := []string{"cat", "cot", "cog", "dog", "dot"}
	solver := NewGreedySolver()
	wordsChains, err := solver.FindWordChains("cat", "dog", wordsList)
	if err != nil {
		panic(err)
	}
	fmt.Println(wordsChains)
	// Output :
	// [[cat cot cog dog] [cat cot dot dog]]
}

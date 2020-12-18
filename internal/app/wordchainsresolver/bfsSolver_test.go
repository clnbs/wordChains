package wordchainsresolver

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFSQueue_Add(t *testing.T) {
	bfsQueue := &BFSQueue{}
	expected := []*BFSWordTreeNode{
		{
			Word:            "test01",
			PreviousElement: nil,
			NextElements:    nil,
		}, {
			Word:            "test02",
			PreviousElement: nil,
			NextElements:    nil,
		},
	}
	bfsQueue.Add(expected[0])
	assert.Equal(t, expected[0], bfsQueue.words[0])
	bfsQueue.Add(expected[1])
	assert.Equal(t, expected[1], bfsQueue.words[1])
}

func TestBFSQueue_Pop(t *testing.T) {
	bfsQueue := &BFSQueue{}
	expected := []*BFSWordTreeNode{
		{
			Word:            "test01",
			PreviousElement: nil,
			NextElements:    nil,
		}, {
			Word:            "test02",
			PreviousElement: nil,
			NextElements:    nil,
		},
		nil,
	}
	bfsQueue.Add(expected[0])
	assert.Equal(t, expected[0], bfsQueue.words[0])
	bfsQueue.Add(expected[1])
	assert.Equal(t, expected[1], bfsQueue.words[1])

	result := bfsQueue.Pop()
	assert.Equal(t, expected[0], result)
	assert.Equal(t, expected[1], bfsQueue.words[0])

	_ = bfsQueue.Pop()

	result = bfsQueue.Pop()
	assert.Equal(t, expected[2], result)
}

func TestBFSQueue_Len(t *testing.T) {
	bfsQueue := &BFSQueue{}
	mockValues := []*BFSWordTreeNode{
		{
			Word:            "test01",
			PreviousElement: nil,
			NextElements:    nil,
		}, {
			Word:            "test02",
			PreviousElement: nil,
			NextElements:    nil,
		},
	}
	bfsQueue.Add(mockValues[0])
	bfsQueue.Add(mockValues[1])
	expected := 2

	assert.Equal(t, expected, bfsQueue.Len())
}

func TestNewBFSWordTreeNode(t *testing.T) {
	headWord := "head"
	nodeWord := "node"
	head := NewBFSWordTreeNode(headWord, nil)
	assert.Equal(t, headWord, head.Word)
	expectedEmptyNode := &BFSWordTreeNode{}

	assert.Equal(t, expectedEmptyNode.PreviousElement, head.PreviousElement)
	node := NewBFSWordTreeNode(nodeWord, head)
	assert.Equal(t, nodeWord, node.Word)
	assert.Equal(t, head, node.PreviousElement)
}

func TestBFSWordTreeNode_Depth(t *testing.T) {
	head := NewBFSWordTreeNode("head", nil)
	node := NewBFSWordTreeNode("node", head)
	secondNode := NewBFSWordTreeNode("n0de", node)
	expectedHeadDepth := 1
	expectedNodeDepth := 2
	expectedSecondNodeDepth := 3

	assert.Equal(t, expectedHeadDepth, head.Depth())
	assert.Equal(t, expectedNodeDepth, node.Depth())
	assert.Equal(t, expectedSecondNodeDepth, secondNode.Depth())
}

func TestBFSWordTreeNode_GetSolution(t *testing.T) {
	expected := []string{"head", "node", "n0de"}
	head := NewBFSWordTreeNode(expected[0], nil)
	node := NewBFSWordTreeNode(expected[1], head)
	secondNode := NewBFSWordTreeNode(expected[2], node)

	result := secondNode.GetSolution()
	assert.Equal(t, expected, result)
}

func TestNewBFSSolverWithParams(t *testing.T) {
	expectedWordList := []string{"cat", "cog", "cot", "dog", "dot"}
	expectedFromWord := "cat"
	expectedToWord := "dog"
	expectedDiscovered := make(map[*BFSWordTreeNode]interface{})
	bfs := NewBFSSolverWithParams(expectedFromWord, expectedToWord, expectedWordList)

	assert.Equal(t, expectedWordList, bfs.wordsList)
	assert.Equal(t, expectedFromWord, bfs.from)
	assert.Equal(t, expectedToWord, bfs.to)
	assert.Equal(t, int(^uint(0)>>1), bfs.bestSolutionDepth)
	assert.Equal(t, expectedDiscovered, bfs.discovered)
}

func TestNewBFSSolver(t *testing.T) {
	expectedDiscovered := make(map[*BFSWordTreeNode]interface{})
	bfs := NewBFSSolver()

	assert.Equal(t, int(^uint(0)>>1), bfs.bestSolutionDepth)
	assert.Equal(t, expectedDiscovered, bfs.discovered)
}

func TestBFSSolver_getUsefulWordOnly(t *testing.T) {
	wordList := []string{"cat", "cog", "cot", "dog", "dot", "code"}
	fromWord := "cat"
	toWord := "dog"
	expectedUsefulWordList := []string{"cog", "cot", "dog", "dot"}

	bfs := NewBFSSolverWithParams(fromWord, toWord, wordList)
	bfs.getUsefulWordOnly()

	assert.Equal(t, expectedUsefulWordList, bfs.usefulWords)
}

type ListPossibleNextNextWordsTestCase struct {
	input    string
	expected []string
}

func TestBFSSolver_listPossibleNextWords(t *testing.T) {
	wordList := []string{"cat", "cog", "cot", "dog", "dot"}
	fromWord := "cat"
	toWord := "dog"
	testCases := []ListPossibleNextNextWordsTestCase{
		{
			input:    "cat",
			expected: []string{"cot"},
		},
		{
			input:    "cot",
			expected: []string{"cog", "dot"},
		},
		{
			input:    "cog",
			expected: []string{"cot", "dog"},
		},
		{
			input:    "dot",
			expected: []string{"cot", "dog"},
		},
	}

	bfs := NewBFSSolverWithParams(fromWord, toWord, wordList)
	bfs.getUsefulWordOnly()

	for _, test := range testCases {
		assert.Equal(t, test.expected, bfs.listPossibleNextWords(test.input))
	}
}

func TestBFSSolver_solveBFS(t *testing.T) {
	wordList := []string{"cat", "cog", "cot", "dog", "dot"}
	fromWord := "cat"
	toWord := "dog"
	expectedResultCount := 2
	expectedSolutions := [][]string{{"cat", "cot", "cog", "dog"}, {"cat", "cot", "dot", "dog"}}

	bfs := NewBFSSolverWithParams(fromWord, toWord, wordList)
	bfs.getUsefulWordOnly()
	bfs.wordTree = NewBFSWordTreeNode(fromWord, nil)
	bfs.solveBFS()

	assert.Equal(t, expectedResultCount, len(bfs.solutions))

	var solutions [][]string
	for _, node := range bfs.solutions {
		solutions = append(solutions, node.GetSolution())
	}
	assert.Equal(t, expectedSolutions, solutions)
}

func TestBFSSolver_FindWordChains(t *testing.T) {
	solver := NewBFSSolver()
	factory := NewFileLoaderFactory(os.Getenv("GOPATH") + "/src/github.com/clnbs/wordChains/assets/app/small_en.txt")
	GeneralWordChainsResolverTest(solver, factory, t)
	solver = NewBFSSolver()
	_, err := solver.FindWordChains("dummy", "to", []string{})
	assert.NotNil(t, err)
}

package wordchainsresolver

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAStarNode(t *testing.T) {
	headWord := "head"
	nodeWord := "node"
	expectedEmptyNode := &AStarNode{}
	head := NewAStarNode(headWord, nil)
	assert.Equal(t, headWord, head.word)
	assert.Equal(t, expectedEmptyNode.previous, head.previous)

	node := NewAStarNode(nodeWord, head)
	assert.Equal(t, nodeWord, node.word)
	assert.Equal(t, head, node.previous)
}

func TestAStarNode_Depth(t *testing.T) {
	head := NewAStarNode("head", nil)
	node := NewAStarNode("node", head)
	deeperNode := NewAStarNode("n0de", node)

	assert.Equal(t, 1, head.Depth())
	assert.Equal(t, 2, node.Depth())
	assert.Equal(t, 3, deeperNode.Depth())
}

func TestAStarNode_GetSolution(t *testing.T) {
	expected := []string{"head", "node", "n0de"}
	head := NewAStarNode("head", nil)
	node := NewAStarNode("node", head)
	deeperNode := NewAStarNode("n0de", node)

	assert.Equal(t, expected, deeperNode.GetSolution())
	assert.Equal(t, expected[:2], node.GetSolution())
	assert.Equal(t, expected[:1], head.GetSolution())
}

func TestAStarSolver_FindWordChains(t *testing.T) {
	solver := NewAStarSolver()
	factory := NewFileLoaderFactory(os.Getenv("GOPATH") + "/src/github.com/clnbs/wordChains/assets/app/small_en.txt")
	expectedResult := [][]string{{"cat", "cot", "cog", "dog"}}
	// We can not use GeneralWordChainsResolverTest for A* because it returns
	// only one expected at each time

	wcr := NewWordChainsResolver(solver, factory)
	err := wcr.LoadDB()
	assert.Nil(t, err)

	result, err := wcr.Solve("cat", "dog")
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)

	solver = NewAStarSolver()
	_, err = solver.FindWordChains("dummy", "to", []string{})
	assert.NotNil(t, err)
}

func TestAStarSolver_GetCurrentBestNode(t *testing.T) {
	aStar := NewAStarSolver()
	aStar.nodeFScore = make(map[*AStarNode]int)
	aStar.openSet = make(map[*AStarNode]interface{})

	mockNode1 := NewAStarNode("fake", nil)
	mockNode2 := NewAStarNode("mock", nil)

	aStar.nodeFScore[mockNode1] = 10
	aStar.nodeFScore[mockNode2] = 3

	aStar.openSet[mockNode1] = nil
	assert.Equal(t, mockNode1, aStar.getCurrentBestNode())

	aStar.openSet[mockNode2] = nil
	assert.Equal(t, mockNode2, aStar.getCurrentBestNode())
}

func TestAStarSolver_helpers(t *testing.T) {
	expectedUsefulWordList := []string{"cot", "cog", "dog", "dot"}
	aStar := NewAStarSolver()
	aStar.wordList = []string{"cat", "cot", "cog", "dog", "dot", "parrot"}
	aStar.from = "cat"
	aStar.to = "dog"
	aStar.getUsefulWordsOnly()

	assert.Equal(t, expectedUsefulWordList, aStar.usefulWords)

	head := NewAStarNode("cat", nil)

	neighbors := aStar.createNeighbors(head)
	assert.Equal(t, 1, len(neighbors))
	assert.Equal(t, "cot", neighbors[0].word)
	assert.Equal(t, head, neighbors[0].previous)
	assert.Equal(t, 2, aStar.getScoreFromGoal(neighbors[0]))
}

func ExampleAStarSolver_FindWordChains() {
	wordsList := []string{"cat", "cot", "cog", "dog", "dot"}
	solver := NewAStarSolver()
	wordChain, err := solver.FindWordChains("cat", "dog", wordsList)
	if err != nil {
		panic(err)
	}
	fmt.Println(wordChain)
	// Output :
	// [[cat cot dot dog]]
}

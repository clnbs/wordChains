package wordchainsresolver

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAStarSolver_FindWordChains(t *testing.T) {
	solver := NewAStarSolver()
	factory := NewFileLoaderFactory(os.Getenv("GOPATH") + "/src/github.com/clnbs/wordChains/assets/app/small_en.txt")

	GeneralWordChainsResolverTest(solver, factory, t)
	solver = NewAStarSolver()
	_, err := solver.FindWordChains("dummy", "to", []string{})
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
	assert.Equal(t, 1, aStar.getScoreFromGoal(neighbors[0]))
}

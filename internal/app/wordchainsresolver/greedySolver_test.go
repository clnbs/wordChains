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

func buildMockWordsTree() *GreedyWordTreeNode {
	head := NewGreedyWordTreeElement("cat", 0, nil)
	head.NextElements = append(head.NextElements, NewGreedyWordTreeElement("cot", 1, head))
	head.NextElements[0].NextElements = append(head.NextElements[0].NextElements,
		NewGreedyWordTreeElement("cog", 2, head.NextElements[0]),
		NewGreedyWordTreeElement("dot", 2, head.NextElements[0]),
	)
	head.NextElements[0].NextElements[0].NextElements = append(head.NextElements[0].NextElements[0].NextElements,
		NewGreedyWordTreeElement("dog", 3, head.NextElements[0].NextElements[0]))
	head.NextElements[0].NextElements[1].NextElements = append(head.NextElements[0].NextElements[1].NextElements,
		NewGreedyWordTreeElement("dog", 3, head.NextElements[0].NextElements[1]))
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

func TestGetNodeDepth(t *testing.T) {
	head := NewGreedyWordTreeElement("test", 0, nil)
	node := NewGreedyWordTreeElement("test_depth_2", 0, head)
	expectedHeadDepth := 1
	expectedNodeDepth := 2
	assert.Equal(t, expectedHeadDepth, head.getNodeDepth())
	assert.Equal(t, expectedNodeDepth, node.getNodeDepth())
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
	head := NewGreedyWordTreeElement("cat", 0, nil)
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

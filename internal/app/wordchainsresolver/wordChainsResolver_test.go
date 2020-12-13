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
	return nil, errors.New("from mock : I am a bad factory, Muahahahaha!")
}

type MockSolver struct {
}

func (solver *MockSolver) FindWordChains(string, string, []string) ([]string, error) {
	return []string{"cat", "cot", "cog", "dog"}, nil
}

func GeneralWordChainsResolverTest(solver Solver, factory Factory, t *testing.T) {
	expectedResult := []string{"cat", "cot", "cog", "dog"}
	wcr := NewWordChainsResolver(solver, factory)
	err := wcr.LoadDB()
	assert.Nil(t, err)
	result, err := wcr.Solve("cat", "dog")
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
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

func TestWordChainsResolver_Solve(t *testing.T) {
	GeneralWordChainsResolverTest(&MockSolver{}, &MockFactory{}, t)
}

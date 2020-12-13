package wordchainsresolver

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockFactory struct {
}

func (factory *MockFactory) LoadDB() ([]string, error) {
	return []string{"cat", "cot", "cog", "dog"}, nil
}

type MockBadFactory struct {
}

func (badFactory *MockBadFactory) LoadDB() ([]string, error) {
	return nil, errors.New("I am a bad factory, Muahahahaha!")
}

type MockAlgorithm struct {
}

func (algo *MockAlgorithm) FindWordChains(string, string, []string) ([]string, error) {
	return []string{"cat", "cot", "cog", "dog"}, nil
}

func GeneralWordChainsResolverTest(algo Algorithm, factory Factory, t *testing.T) {
	expectedResult := []string{"cat", "cot", "cog", "dog"}
	wcr := NewWordChainsResolver(algo, factory)
	err := wcr.LoadDB()
	assert.Nil(t, err)
	result, err := wcr.Solve("cat", "dog")
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestNewWordChainsResolver(t *testing.T) {
	algo := &MockAlgorithm{}
	factory := &MockFactory{}
	wcr := NewWordChainsResolver(algo, factory)
	assert.Equal(t, wcr.algorithm, algo)
	assert.Equal(t, wcr.factory, factory)
}

func TestWordChainsResolver_LoadDB(t *testing.T) {
	expectedResult := []string{"cat", "cot", "cog", "dog"}
	wcr := NewWordChainsResolver(&MockAlgorithm{}, &MockFactory{})
	err := wcr.LoadDB()
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, wcr.wordList)

	wcr = NewWordChainsResolver(&MockAlgorithm{}, &MockBadFactory{})
	err = wcr.LoadDB()
	assert.NotNil(t, err)
}

func TestWordChainsResolver_Solve(t *testing.T) {
	GeneralWordChainsResolverTest(&MockAlgorithm{}, &MockFactory{}, t)
}

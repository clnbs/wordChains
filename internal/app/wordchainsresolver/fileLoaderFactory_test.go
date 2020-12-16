package wordchainsresolver

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileLoaderFactory_LoadDB(t *testing.T) {
	flf := NewFileLoaderFactory(os.Getenv("GOPATH") + "/src/github.com/clnbs/wordChains/assets/app/small_en.txt")
	wordList, err := flf.LoadDB()
	assert.Nil(t, err)
	assert.Equal(t, 58110, len(wordList))

	flf = NewFileLoaderFactory("/badpath/thing.txt")
	_, err = flf.LoadDB()
	assert.NotNil(t, err)
}

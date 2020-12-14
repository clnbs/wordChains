package wordchainsresolver

import (
	"bufio"
	"os"
)

//FileLoaderFactory struct implements Factory interface
type FileLoaderFactory struct {
	path string
}

// NewFileLoaderFactory is a FileLoaderFactory constructor
func NewFileLoaderFactory(path string) *FileLoaderFactory {
	return &FileLoaderFactory{path: path}
}

// LoadDB implement Factory interface. It read a file containing a word per line
func (fileLoader *FileLoaderFactory) LoadDB() ([]string, error) {
	var wordList []string
	file, err := os.Open(fileLoader.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return wordList, nil
}

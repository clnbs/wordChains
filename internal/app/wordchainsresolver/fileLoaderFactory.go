package wordchainsresolver

import (
	"bufio"
	"os"
)

type FileLoaderFactory struct {
	path string
}

func NewFileLoaderFactory(path string) *FileLoaderFactory {
	return &FileLoaderFactory{path: path}
}

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

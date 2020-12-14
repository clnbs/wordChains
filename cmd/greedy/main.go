package main

import (
	"fmt"
	"github.com/clnbs/wordChains/internal/app/wordchainsresolver"
	"os"
)

func main() {
	solver := wordchainsresolver.NewGreedySolver()
	factory := wordchainsresolver.NewFileLoaderFactory(os.Getenv("GOPATH") + "/src/github.com/clnbs/wordChains/assets/app/wordlist.txt")
	wcr := wordchainsresolver.NewWordChainsResolver(solver, factory)
	err := wcr.LoadDB()
	if err != nil {
		panic(err)
	}
	path, err := wcr.Solve("ruby", "code")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	path, err = wcr.Solve("lead", "gold")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	path, err = wcr.Solve("gold", "lead")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	path, err = wcr.Solve("cat", "dog")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	path, err = wcr.Solve("cake", "gold")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	path, err = wcr.Solve("bar", "oil")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	path, err = wcr.Solve("oil", "bar")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	path, err = wcr.Solve("abates", "hollow")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
}

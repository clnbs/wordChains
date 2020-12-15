package main

import (
	"fmt"
	"os"

	"github.com/clnbs/wordChains/internal/app/wordchainsresolver"
)

func usage(programName string) {
	fmt.Println("usage :\t\t", programName, "path/to/wordlist.txt word1 word2")
	fmt.Println("example :\t", programName, "./assets/app/wordlist.txt cat dog")
}

func printSolutions(solutions [][]string) {
	if len(solutions) == 0 {
		fmt.Println("no solution found")
		return
	}
	fmt.Println("found", len(solutions), "solution(s)")
	for index, chain := range solutions {
		fmt.Print("solution #", index+1, " : ")
		for chainIndex, word := range chain {
			fmt.Print(word)
			if chainIndex == len(chain)-1 {
				continue
			}
			fmt.Print(" -> ")
		}
		fmt.Println()
	}
}

func main() {
	programName, args := os.Args[0], os.Args[1:]
	if len(args) != 3 {
		usage(programName)
		return
	}
	solver := wordchainsresolver.NewBFSSolver()
	factory := wordchainsresolver.NewFileLoaderFactory(args[0])
	wcr := wordchainsresolver.NewWordChainsResolver(solver, factory)
	err := wcr.LoadDB()
	if err != nil {
		fmt.Println("error while loading word list :", err)
		return
	}
	path, err := wcr.Solve(args[1], args[2])
	if err != nil {
		fmt.Println("error while solving word chains :", err)
		return
	}
	printSolutions(path)
}

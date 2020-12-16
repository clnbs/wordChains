package main

import (
	"fmt"
	"github.com/clnbs/wordChains/internal/app/wordchainsresolver"
	"os"
	"strings"
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
	filePath := args[0]
	word1 := args[1]
	word2 := args[2]
	solver := wordchainsresolver.NewGreedySolver()
	factory := wordchainsresolver.NewFileLoaderFactory(filePath)
	wcr := wordchainsresolver.NewWordChainsResolver(solver, factory)
	err := wcr.LoadDB()
	if err != nil {
		fmt.Println("error while loading word list :", err)
		return
	}
	if !wcr.IsWordInDB(word1) {
		fmt.Println(word1, "is not in your database")
		return
	}
	if !wcr.IsWordInDB(args[2]) {
		fmt.Println(args[2], "is not in your database")
		return
	}
	fmt.Println("looking for word chains from", word1, "to", word2+", please wait ...")
	path, err := wcr.Solve(strings.ToLower(word1), strings.ToLower(word2))
	if err != nil {
		fmt.Println("error while solving word chains :", err)
		return
	}
	printSolutions(path)
}

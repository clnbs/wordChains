package wordchainsresolver

// Factory handle everything linked to loading data
type Factory interface {
	LoadDB() ([]string, error)
}

// Solver handle calculus part of the word chains problem
type Solver interface {
	FindWordChains(string, string, []string) ([]string, error)
}

// WordChainsResolver wrap Solver and Factory interfaces by holding
// the word lis to process
type WordChainsResolver struct {
	solver   Solver
	factory  Factory
	wordList []string
}

// NewWordChainsResolver WordChainsResolver struct constructor
func NewWordChainsResolver(solver Solver, factory Factory) *WordChainsResolver {
	return &WordChainsResolver{solver: solver, factory: factory}
}

// LoadDB Factory wrapper
func (wcr *WordChainsResolver) LoadDB() error {
	var err error
	wcr.wordList, err = wcr.factory.LoadDB()
	if err != nil {
		return err
	}
	return nil
}

// Solve Solver wrapper
func (wcr *WordChainsResolver) Solve(from, to string) ([]string, error) {
	return wcr.solver.FindWordChains(from, to, wcr.wordList)
}

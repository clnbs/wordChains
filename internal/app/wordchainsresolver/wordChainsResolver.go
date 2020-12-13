package wordchainsresolver

// Factory handle everything linked to loading data
type Factory interface {
	LoadDB() ([]string, error)
}

// Algorithm handle calculus part of the word chains problem
type Algorithm interface {
	FindWordChains(string, string, []string) ([]string, error)
}

// WordChainsResolver wrap Algorithm and Factory interfaces by holding
// the word lis to process
type WordChainsResolver struct {
	algorithm Algorithm
	factory   Factory
	wordList  []string
}

// NewWordChainsResolver WordChainsResolver struct constructor
func NewWordChainsResolver(algo Algorithm, factory Factory) *WordChainsResolver {
	return &WordChainsResolver{algorithm: algo, factory: factory}
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

// Solve Algorithm wrapper
func (wcr *WordChainsResolver) Solve(from, to string) ([]string, error) {
	return wcr.algorithm.FindWordChains(from, to, wcr.wordList)
}

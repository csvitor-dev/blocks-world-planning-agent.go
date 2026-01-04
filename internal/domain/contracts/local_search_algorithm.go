package contracts

type LocalSearchAlgorithm interface {
	Execute() ([]string, error)
}

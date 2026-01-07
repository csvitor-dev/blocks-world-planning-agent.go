package contracts

type LocalSearchAlgorithm interface {
	Execute() ([]string, int, int)
	Solution(goal BlocksWorldState) []string
	Name() string
}

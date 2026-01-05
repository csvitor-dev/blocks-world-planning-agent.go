package factories

import (
	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/domain/contracts"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/src/algorithms"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/types"
)

var (
	mapAlgorithms = map[string]types.AlgorithmConstructor{
		"BFS": func(p contracts.PlanningContract) contracts.LocalSearchAlgorithm { return algorithms.NewBFS(p) },
	}
)

func MakeAlgorithm(algorithm string) types.AlgorithmConstructor {
	return mapAlgorithms[algorithm]
}

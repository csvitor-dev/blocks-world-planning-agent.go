package types

import (
	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/domain/contracts"
)

type AlgorithmConstructor func(p contracts.PlanningContract) contracts.LocalSearchAlgorithm

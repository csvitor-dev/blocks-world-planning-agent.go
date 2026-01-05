package algorithms

import (
	"fmt"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/domain/contracts"
)

type BFS struct {
	planning contracts.PlanningContract
}

func NewBFS(planning contracts.PlanningContract) *BFS {
	return &BFS{planning: planning}
}

func (b *BFS) Execute() ([]string, error) {
	fmt.Println("Executing BFS...")
	return nil, nil
}

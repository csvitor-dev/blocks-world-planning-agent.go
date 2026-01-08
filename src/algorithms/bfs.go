package algorithms

import (
	"slices"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/domain/contracts"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/sets"
	pkg_slices "github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/slices"
)

type BFS struct {
	planning       contracts.PlanningContract
	explored       sets.Set[string]
	frontier       []contracts.BlocksWorldState
	generatedNodes int
}

func NewBFS(planning contracts.PlanningContract) *BFS {
	return &BFS{
		planning: planning,
		explored: sets.NewSet[string](),
		frontier: []contracts.BlocksWorldState{planning.CurrentState()},
	}
}

func (bfs *BFS) Execute() ([]string, int, int) {
	for len(bfs.frontier) > 0 {
		var state contracts.BlocksWorldState
		bfs.frontier, state, _ = pkg_slices.PopLeft(bfs.frontier)

		if bfs.explored.Has(state.Key()) {
			continue
		}

		if bfs.planning.IsGoalState(state) {
			return bfs.Solution(state), bfs.generatedNodes, len(bfs.explored)
		}
		bfs.explored.Add(state.Key())

		for _, successor := range state.Successors(bfs.planning.Actions()) {
			bfs.generatedNodes += 1

			if !bfs.explored.Has(successor.Key()) {
				bfs.frontier = append(bfs.frontier, successor)
			}
		}
	}
	return nil, bfs.generatedNodes, len(bfs.explored)
}

func (bfs *BFS) Name() string {
	return "Breadth-First Search"
}

func (bfs *BFS) Solution(goal contracts.BlocksWorldState) []string {
	solution := []string{}
	hook := goal

	for hook.Parent() != nil {
		solution = append(solution, hook.Step())
		hook = hook.Parent()
	}
	slices.Reverse(solution)

	return solution
}

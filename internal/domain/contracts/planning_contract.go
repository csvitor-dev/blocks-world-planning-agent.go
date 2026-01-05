package contracts

import (
	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/types"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/sets"
)

type PlanningContract interface {
	CurrentState() BlocksWorldState
	States() map[string]sets.Set[int]
	Actions() map[string]types.Action
	Remap(state sets.Set[int]) sets.Set[string]
	Solution(goal BlocksWorldState) []string
	SetGoal(goal sets.Set[int])
	SetInitial(initial sets.Set[int])
	Copy() PlanningContract
}

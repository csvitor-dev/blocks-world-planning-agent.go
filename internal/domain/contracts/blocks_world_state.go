package contracts

import (
	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/types"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/sets"
)

type BlocksWorldState interface {
	Step() string
	Current() sets.Set[int]
	Successors(actions map[string]types.Action) []BlocksWorldState
	Parent() BlocksWorldState
	String() string
	Key() string
	Equals(other BlocksWorldState) bool
	LessThan(other BlocksWorldState) bool
	GreaterThan(other BlocksWorldState) bool
	Hash() int
}

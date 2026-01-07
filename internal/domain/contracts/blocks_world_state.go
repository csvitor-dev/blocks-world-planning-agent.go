package contracts

import "github.com/csvitor-dev/blocks-world-planning-agent.go/internal/types"

type BlocksWorldState interface {
	Successors(actions map[string]types.Action) []BlocksWorldState
	String() string
	Key() string
	Equals(other BlocksWorldState) bool
	LessThan(other BlocksWorldState) bool
	GreaterThan(other BlocksWorldState) bool
	Hash() int
}

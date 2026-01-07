package domain

import (
	"errors"
	"fmt"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/domain/contracts"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/types"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/sets"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/utils"
)

type BlocksWorldState struct {
	Id               string
	Current          sets.Set[int]
	AvailableActions map[string]types.Action
	Parent           *BlocksWorldState
	G                int
	H                int
	F                int
}

func NewBlocksWorldState(current sets.Set[int], actions map[string]types.Action, name string, parent *BlocksWorldState, realCost ...int) (*BlocksWorldState, error) {
	if len(realCost) > 1 {
		return nil, errors.New("blocks_world_state: only one 'realCost' argument is allowed")
	}
	var cost int = 0

	if len(realCost) == 1 {
		cost = realCost[0]
	}
	state := &BlocksWorldState{
		Id:      name,
		Current: current,
		Parent:  parent,
		G:       cost,
		H:       0,
		F:       cost,
	}
	state.AvailableActions = state.filterAvailableActions(actions)
	return state, nil
}

func (s *BlocksWorldState) Successors(actions map[string]types.Action) []contracts.BlocksWorldState {
	var out []contracts.BlocksWorldState

	for name, action := range s.AvailableActions {
		newState := s.expand(name, action, actions)
		out = append(out, newState)
	}
	return out
}

func (s *BlocksWorldState) expand(actionName string, action types.Action, actions map[string]types.Action) *BlocksWorldState {
	transitionState := s.Current.Difference(action.Pre)
	newState := resolveConsistentState(transitionState, action.Post)

	st, _ := NewBlocksWorldState(newState, actions, actionName, s)
	return st
}

func (s *BlocksWorldState) filterAvailableActions(actions map[string]types.Action) map[string]types.Action {
	out := make(map[string]types.Action)
	for name, cond := range actions {
		if cond.Pre.IsSubsetOf(s.Current) {
			out[name] = cond
		}
	}
	return out
}

func resolveConsistentState(transition sets.Set[int], post sets.Set[int]) sets.Set[int] {
	positive := make(sets.Set[int])

	for v := range post {
		if v > 0 {
			positive[v] = struct{}{}
		}
	}
	return transition.Union(positive)
}

func (s *BlocksWorldState) Key() string {
	return sets.SortedString(s.Current)
}

func (s *BlocksWorldState) String() string {
	return fmt.Sprintf("State(%s)", s.Key())
}

func (s *BlocksWorldState) Equals(other contracts.BlocksWorldState) bool {
	return s.Key() == other.Key()
}

func (s *BlocksWorldState) LessThan(other contracts.BlocksWorldState) bool {
	return s.Key() == other.Key()
}

func (s *BlocksWorldState) GreaterThan(other contracts.BlocksWorldState) bool {
	return s.Key() == other.Key()
}

func (s *BlocksWorldState) Hash() int {
	return int(utils.Hash(s.Key()))
}

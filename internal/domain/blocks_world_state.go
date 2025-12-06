package domain

import (
	"errors"
	"fmt"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/sets"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/types"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/utils"
)

type BlocksWorldState struct {
	Id               string
	Current          sets.Set[int]
	Key              string
	AvailableActions map[string]types.Action
	Parent           *BlocksWorldState
}

func NewBlocksWorldState(current sets.Set[int], actions map[string]types.Action, name string, parent *BlocksWorldState) (*BlocksWorldState, error) {
	if !isValidState(current) {
		return nil, errors.New("blocks_world_state: invalid current state")
	}

	state := &BlocksWorldState{
		Id:      name,
		Current: current,
		Key:     sets.SortedString(current),
		Parent:  parent,
	}
	state.AvailableActions = state.filterAvailableActions(actions)
	return state, nil
}

func (s *BlocksWorldState) Successors(actions map[string]types.Action) []*BlocksWorldState {
	var out []*BlocksWorldState

	for name, action := range s.AvailableActions {
		newState := s.expand(name, action, actions)
		out = append(out, newState)
	}

	return out
}

func (s *BlocksWorldState) expand(actionName string, action types.Action, actions map[string]types.Action) *BlocksWorldState {
	transitionState := subtract(s.Current, action.Pre)
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

func subtract(a, b sets.Set[int]) sets.Set[int] {
	out := make(sets.Set[int])
	for v := range a {
		if _, ok := b[v]; !ok {
			out[v] = struct{}{}
		}
	}
	return out
}

func isValidState(state sets.Set[int]) bool {
	absSet := make(map[int]struct{})
	for v := range state {
		absV := v
		if v < 0 {
			absV = -v
		}
		if _, ok := absSet[absV]; ok {
			return false
		}
		absSet[absV] = struct{}{}
	}
	return true
}

func (s *BlocksWorldState) String() string {
	return fmt.Sprintf("State(%s)", s.Key)
}

func (s *BlocksWorldState) Equals(o *BlocksWorldState) bool {
	return s.Key == o.Key
}

func (s *BlocksWorldState) Hash() int {
	return int(utils.Hash(s.Key))
}

package domain

import (
	"math"
	"strings"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/domain/contracts"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/types"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/sets"
)

type Planning struct {
	strips       StripsNotation
	_map         map[string]int
	inverseMap   map[int]string
	actions      map[string]types.Action
	planner      contracts.LocalSearchAlgorithm
	initialState sets.Set[int]
	goalState    sets.Set[int]
	stateSpace   *BlocksWorldState
}

func NewPlanning(strips StripsNotation, instanceId string) (*Planning, error){
	_map := mapFacts(strips)
	inverseMap := inverseMapFacts(_map)
	current := resolveFacts(strips.InitialState, _map)
	actions := resolveActions(strips.Actions, _map)
	stateNode, err := NewBlocksWorldState(current, actions, "root-"+instanceId, nil)

	if err != nil {
		return nil, err
	}

	return &Planning{
		strips:     strips,
		_map:       _map,
		inverseMap: inverseMap,
		actions:    actions,
		initialState: current,
		stateSpace: stateNode,
	}, nil
}

func (p *Planning) Remap(state sets.Set[int]) sets.Set[string] {
	remapped := sets.NewSet[string]()

	for fact := range state {
		remapped.Add(p.inverseMap[fact])
	}
	return remapped
}

func mapFacts(strips StripsNotation) map[string]int {
	hook := make(map[string]int)

	for fact := range strips.AvaliableFacts() {
		var negativeFact string

		if strings.HasPrefix(fact, "~") {
			negativeFact = fact[1:]
		} else {
			negativeFact = "~" + fact
		}

		if _, exists := hook[fact]; !exists {
			existentValue := float64(hook[negativeFact])

			if !strings.HasPrefix(fact, "~") {
				if existentValue == 0 {
					hook[fact] = len(hook) + 1
				} else {
					hook[fact] = int(math.Abs(existentValue))
				}
			} else {
				if existentValue == 0 {
					hook[fact] = -(len(hook) + 1)
				} else {
					hook[fact] = -int(existentValue)
				}
			}
		}
	}
	return hook
}

func inverseMapFacts(factMap map[string]int) map[int]string {
	inverse := make(map[int]string)

	for fact, value := range factMap {
		inverse[value] = fact
	}
	return inverse
}

func resolveFacts(facts []string, factMap map[string]int) sets.Set[int] {
	resolved := sets.NewSet[int]()

	for _, fact := range facts {
		resolved.Add(factMap[fact])
	}
	return resolved
}

func resolveActions(actions map[string][2][]string, factMap map[string]int) map[string]types.Action {
	actionHook := make(map[string]types.Action)

	for actionName, conditions := range actions {
		actionHook[actionName] = types.Action{
			Pre:  resolveFacts(conditions[0], factMap),
			Post: resolveFacts(conditions[1], factMap),
		}
	}
	return actionHook
}

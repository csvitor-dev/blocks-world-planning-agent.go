package domain

import (
	"slices"
	"strings"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/sets"
)

type StripsNotation struct {
	facts        sets.Set[string]
	Actions      map[string][2][]string
	InitialState []string
	GoalState    []string
}

func NewStrips(actionsSet []string, initial string, goal string) StripsNotation {
	actions, facts := createActionsWithFacts(actionsSet)

	return StripsNotation{
		facts:        facts,
		Actions:      actions,
		InitialState: splitFacts(initial),
		GoalState:    splitFacts(goal),
	}
}

func createActionsWithFacts(rawActions []string) (map[string][2][]string, sets.Set[string]) {
	actionHook := map[string][2][]string{}
	factHook := sets.NewSet[string]()

	for i := 0; i < len(rawActions); i += 3 {
		preconditions, postConditions := splitFacts(rawActions[i+1]), splitFacts(rawActions[i+2])
		factHook.AddFrom(slices.Concat(preconditions, postConditions))

		actionHook[rawActions[i]] = [...][]string{preconditions, postConditions}
	}
	return actionHook, factHook
}

func splitFacts(rawFacts string) []string {
	return strings.Split(rawFacts, ";")
}

func (s StripsNotation) AvaliableFacts() sets.Set[string] {
	return s.facts
}

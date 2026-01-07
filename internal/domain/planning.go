package domain

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"strings"
	"time"
	"unsafe"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/domain/contracts"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/types"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/sets"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/src/support/factories"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/utils"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/utils/report"
)

type Planning struct {
	strips       StripsNotation
	instanceId   string
	_map         map[string]int
	inverseMap   map[int]string
	actions      map[string]types.Action
	planner      contracts.LocalSearchAlgorithm
	initialState sets.Set[int]
	goalState    sets.Set[int]
	stateSpace   *BlocksWorldState
	showReport   bool
	sizeOf       uintptr
}

func NewPlanning(strips StripsNotation, instanceId string) (*Planning, error) {
	_map := mapFacts(strips)
	inverseMap := inverseMapFacts(_map)
	current := resolveFacts(strips.InitialState, _map)
	actions := resolveActions(strips.Actions, _map)
	state, err := NewBlocksWorldState(current, actions, "root-"+instanceId, nil)

	if err != nil {
		return nil, err
	}

	return &Planning{
		strips:       strips,
		instanceId:   instanceId,
		_map:         _map,
		inverseMap:   inverseMap,
		actions:      actions,
		initialState: current,
		stateSpace:   state,
		showReport:   true,
		sizeOf:       unsafe.Sizeof(state),
	}, nil
}

func (p *Planning) Remap(state sets.Set[int]) sets.Set[string] {
	remapped := sets.NewSet[string]()

	for fact := range state {
		remapped.Add(p.inverseMap[fact])
	}
	return remapped
}

func (p *Planning) OffReport() {
	p.showReport = false
}

func (p *Planning) Plan() {
	if p.planner == nil {
		log.Fatalln("No planning algorithm set.")
		return
	}
	statsBefore, statsAfter := runtime.MemStats{}, runtime.MemStats{}
	runtime.GC()
	runtime.ReadMemStats(&statsBefore)

	start := time.Now()
	solution, generatedNodes, exploredNodes := p.planner.Execute()
	runtime.ReadMemStats(&statsAfter)
	report := report.NewPlanningReport(statsBefore, statsAfter)
	report.Elapsed = time.Since(start).Seconds()

	log.Println("Planning completed successfully.")

	if p.showReport == false {
		fmt.Printf("Instance: %s\n", p.instanceId)

		if solution != nil {
			fmt.Printf("Solution Found! Steps: %d Time: %.6f s\n", len(solution), report.Elapsed)
		} else {
			fmt.Printf("No solution found. Time: %.6fs\n", report.Elapsed)
		}
		return
	}
	p.report(solution, generatedNodes, exploredNodes, report)
}

func (p *Planning) IsGoalState(state contracts.BlocksWorldState) bool {
	return p.goalState.Equals(goal)
}

func (p *Planning) CurrentState() contracts.BlocksWorldState {
	return p.stateSpace
}

func (p *Planning) States() (sets.Set[int], sets.Set[int]) {
	return p.initialState, p.goalState
}

func (p *Planning) Actions() map[string]types.Action {
	return p.actions
}

func (p *Planning) Copy() contracts.PlanningContract {
	newPlanning, _ := NewPlanning(p.strips, p.instanceId)

	return newPlanning
}

func (p *Planning) SetAlgorithm(algorithm string) error {
	newPlanner := factories.MakeAlgorithm(algorithm)(p)

	if newPlanner == nil {
		return fmt.Errorf("Algorithm %s not found.", algorithm)
	}
	p.planner = newPlanner
	return nil
}

func (p *Planning) SetGoal(goal sets.Set[int]) {
	p.goalState = goal
}

func (p *Planning) SetInitial(initial sets.Set[int]) {
	p.initialState = initial
}

func (p *Planning) report(result []string, expansions int, explorations int, report report.PlanningReport) {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println(utils.Center("Execution summary", 60))
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Algorithm          : %s\n", p.planner.Name())
	fmt.Printf("Instance           : %s\n", p.instanceId)
	fmt.Printf("Time elapsed       : %.6f s\n", report.Elapsed)
	fmt.Printf("Expanded nodes     : %d\n", expansions)
	fmt.Printf("Explored nodes     : %d\n", explorations)
	fmt.Printf("Total memory cost : %.2f KB\n", float64(explorations*int(p.sizeOf))/1024)
	fmt.Printf(
		"Memory usage      : Allocated=%.2f KB; Stack=%.2f KB; Heap=%.2f KB\n",
		report.TotalAllocated, report.TotalStackInUse, report.TotalHeapInUse,
	)
	fmt.Println(strings.Repeat("-", 60))

	if result != nil {
		fmt.Printf("Solution found! Steps: %d\n", len(result))
		for _, step := range result {
			fmt.Println(step)
		}
	} else {
		fmt.Println("No solution found")
	}
	fmt.Println(strings.Repeat("=", 60))
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

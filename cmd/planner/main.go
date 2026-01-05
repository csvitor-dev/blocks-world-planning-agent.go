package main

import (
	"log"
	"time"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/domain"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/src/services/parser"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/src/support/factories"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/utils/cmd"
)

func execute(ref, algorithm string, dontShowOutput bool) {
	stripsInstance := parser.GetInstance(ref)
	planning, err := domain.NewPlanning(stripsInstance, ref)

	if err != nil {
		log.Fatalf("Error while creating planning for instance %s: %s", ref, err)
	}

	if dontShowOutput {
		planning.OffReport()
	}
	planner := factories.MakeAlgorithm(algorithm)

	if planner == nil {
		log.Fatalf("Algorithm %s not found for instance %s", algorithm, ref)
	}
	planning.SetAlgorithm(planner(planning))

	if err != nil {
		log.Fatalf("Error while setting algorithm for instance %s: %s", ref, err)
	}
	err = planning.Plan()

	if err != nil {
		log.Fatalf("Error while planning for instance %s: %s", ref, err)
	}
}

func main() {
	flags, err := cmd.PluckFlagsFromArgs("instance", "algorithm", "only-execute")

	if err != nil {
		log.Fatalln(err)
	}
	instances := flags["instance"].([]string)
	algorithm := flags["algorithm"].(string)
	onlyExecute := flags["only-execute"].(bool)

	if len(instances) > 0 {
		execute(instances[0], algorithm, onlyExecute)
		return
	}

	for _, ref := range instances {
		execute(ref, algorithm, onlyExecute)
		time.Sleep(time.Second * 3)
	}
}

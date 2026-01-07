package main

import (
	"log"
	"time"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/domain"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/src/services/parser"
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
	planning.SetAlgorithm(algorithm)

	if err != nil {
		log.Fatalf("Error while setting algorithm for instance %s: %s", ref, err)
	}
	planning.Plan()
}

func main() {
	flags, err := cmd.PluckFlagsFromArgs("instance", "algorithm", "only-execute")

	if err != nil {
		log.Fatalln(err)
	}
	instances := flags["instance"].([]string)
	algorithm := flags["algorithm"].(string)
	onlyExecute := flags["only-execute"].(bool)

	if len(instances) == 1 {
		execute(instances[0], algorithm, onlyExecute)
		return
	}

	for _, ref := range instances {
		execute(ref, algorithm, onlyExecute)
		time.Sleep(time.Second * 3)
	}
}

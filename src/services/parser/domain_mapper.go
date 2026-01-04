package parser

import (
	"errors"
	"fmt"
	"log"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/domain"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/filesystem"
)

func GetInstance(ref string) domain.StripsNotation {
	file, err := filesystem.ResolvePath("./resources/planningsat", fmt.Sprintf("blocks-%s.strips", ref))

	if err != nil {
		log.Fatalln(err)
	}
	raw, err := filesystem.Read(file)

	if err != nil {
		log.Fatalln(err)
	}
	separationLine := len(raw) - 3
	actionsSet, initialState, goalState := raw[:separationLine-1], raw[separationLine], raw[separationLine+1]

	if len(actionsSet)%3 != 0 {
		log.Fatalln(len(actionsSet), errors.New("domain_mapper: invalid length in STRIPS Notation file"))
	}
	return domain.NewStrips(actionsSet, initialState, goalState)
}

package parser

import (
	"errors"
	"fmt"
	"log"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/filesystem"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/types"
)

func GetInstance(ref string) types.StripsNotation {
	file, err := filesystem.ResolvePath("./resources/planningsat", fmt.Sprintf("blocks-%s.strips", ref))

	if err != nil {
		log.Fatalln(err)
	}
	raw, err := filesystem.Read(file)

	if err != nil {
		log.Fatalln(err)
	}
	separationLine := len(raw) - 3
	actionsSet, states := raw[:separationLine-1], raw[separationLine+1:]

	if len(actionsSet)%3 != 0 {
		log.Fatalln(errors.New("domain_mapper: invalid length in STRIPS Notation file"))
	}
	return types.NewStrips(actionsSet, states[0], states[1])
}

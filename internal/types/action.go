package types

import "github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/sets"

type Action struct {
	Pre  sets.Set[int]
	Post sets.Set[int]
}

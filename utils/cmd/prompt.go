package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/slices"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/utils"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/utils/constants"
)

func PluckFlagsFromArgs(searchFor ...string) (map[string]any, error) {
	args := mapArgs()

	if len(args) == 0 {
		return nil, errors.New("The flags need to be provided.")
	}
	return buildOutputMap(args, searchFor), nil
}

func mapArgs() map[string]any {
	rawArgs := os.Args[1:]
	filteredArgs := slices.Filter(rawArgs, func(arg string, i int) bool {
		return !utils.IsDigit(arg) && strings.Contains(arg, "--")
	})
	mappedArgs := slices.Map(filteredArgs, func(arg string, i int) []string {
		arg = strings.ReplaceAll(arg, "--", "")
		return strings.Split(arg, "=")
	})
	output := make(map[string]any)

	for _, pair := range mappedArgs {
		if len(pair) == 1 {
			output[pair[0]] = true
		} else {
			output[pair[0]] = resolveFlagListPattern(pair[0], pair[1])
		}
	}
	return output
}

func resolveFlagListPattern(flag string, value string) any {
	if flag != "instance" {
		return value
	}
	if constants.UniqueValueConstraint.MatchString(value) {
		return []string{value}
	}
	return strings.Split(value[1:len(value)-1], ",")
}

func buildOutputMap(args map[string]any, searchFor []string) map[string]any {
	output := make(map[string]any)

	for _, target := range searchFor {
		if value, exists := args[target]; exists {
			output[target] = value
		} else {
			output[target] = false
		}
	}
	return output
}

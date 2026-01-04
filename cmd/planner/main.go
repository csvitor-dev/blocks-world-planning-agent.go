package main

import (
	"fmt"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/internal/domain"
	"github.com/csvitor-dev/blocks-world-planning-agent.go/src/services/parser"
)

func main() {
	instance := parser.GetInstance("4-0")
	planning, err := domain.NewPlanning(instance, "4-0")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(planning)
}

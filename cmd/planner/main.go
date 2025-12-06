package main

import (
	"fmt"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/src/services/parser"
)

func main() {
	instance := parser.GetInstance("4-0")

	fmt.Println(instance)
}

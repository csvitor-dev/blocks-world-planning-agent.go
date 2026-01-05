package constants

import "regexp"

var (
	UniqueValueConstraint = regexp.MustCompile(`^\d+-\d+$`)
)

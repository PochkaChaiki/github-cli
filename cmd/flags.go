package cmd

import (
	"fmt"
)

type arrayOfStrings []string

var owner, repo string
var title, text, milestone, stateReason, state, assignee string
var assignees arrayOfStrings
var number int

func (l *arrayOfStrings) String() string {
	return fmt.Sprintf("%v", *l)
}
func (l *arrayOfStrings) Set(value string) error {
	*l = append(*l, value)
	return nil
}
func (l *arrayOfStrings) Type() string {
	return "arrayOfStrings"
}

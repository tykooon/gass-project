package gasscore

import (
	"errors"
)

type Scenario struct {
	Name      string
	Variables []string
	Commands  []string
}

type ScenarioLib map[string]*Scenario

func IsValidCommand(command string) bool {
	_, exists := commands[command]
	return exists
}

func (s *Scenario) Run(params []string) error {
	return errors.New("not implemented")
}

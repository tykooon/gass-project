package gasscore

import (
	"errors"

	"gopkg.in/yaml.v3"
)

type Scenario struct {
	Name      string
	Variables []string `yaml:"variables"`
	Commands  []string `yaml:"commands"`
}

func NewScenario(name string) *Scenario {
	return &Scenario{
		Name:      name,
		Variables: make([]string, 0),
		Commands:  make([]string, 0),
	}
}

func (scenario *Scenario) Text() (text string) {
	yamlData, err := yaml.Marshal(scenario)
	if err == nil {
		text = string(yamlData)
	}
	return
}

func (scenario *Scenario) ParamsCount() int {
	return len(scenario.Variables)
}

func (scenario *Scenario) Clear() {
	scenario.Variables = make([]string, 0)
	scenario.Commands = make([]string, 0)
}

func (scenario *Scenario) AddVariables(names ...string) {
	scenario.Variables = append(scenario.Variables, names...)
}

func (scenario *Scenario) AddCommands(names ...string) {
	scenario.Commands = append(scenario.Commands, names...)
}

func (s *Scenario) Run(params []string) (string, error) {
	return "", errors.New("not implemented")
}

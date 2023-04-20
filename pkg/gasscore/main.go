package gasscore

type Scenario struct {
	Name      string
	Variables []string
	Commands  []string
}

type ScenarioLib map[string]Scenario

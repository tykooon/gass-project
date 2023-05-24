package gasscore

type ScenarioLib map[string]*Scenario

func (scLib ScenarioLib) Len() int {
	return len(scLib)
}

func (scLib ScenarioLib) ClearAll() {
	for _, v := range scLib {
		v.Clear()
	}
}

func (scLib ScenarioLib) Delete(name string) {
	delete(scLib, name)

}

func (scLib ScenarioLib) DeleteAll(name string) {
	for k := range scLib {
		delete(scLib, k)
	}
}

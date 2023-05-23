package gasscore

import "errors"

type GassContext struct {
	command  string
	scenario string
	lib      ScenarioLib
	params   []string
}

func NewContext(args []string, lib ScenarioLib) (*GassContext, error) {
	result := &GassContext{
		lib:      lib,
		command:  "",
		scenario: "",
		params:   make([]string, 0),
	}
	offset := 0
	if len(args) != 0 {
		if IsValidCommand(args[0]) {
			result.command = args[0]
			offset++
		}
	}

	if offset < len(args) {
		if _, ok := lib[args[offset]]; ok {
			result.scenario = args[offset]
			offset++
		}
	}

	if offset == 0 {
		return nil, errors.New("no commands or scenarios were found")
	}

	if offset < len(args) {
		result.params = append(result.params, args[offset:]...)
	}

	return result, nil
}

func (ctx *GassContext) CommandExec() error {
	if ctx.command == "" {
		return ctx.lib[ctx.scenario].Run(ctx.params)
	}
	return commands[ctx.command](ctx.lib[ctx.scenario].Name, ctx.params)
}

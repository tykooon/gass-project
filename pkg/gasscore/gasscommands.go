package gasscore

import (
	"errors"
	"fmt"
)

var commands = map[string]func(*GassContext) (string, error){
	"clear-all":  clearAllCommand,
	"clear":      clearCommand,
	"help":       helpCommand,
	"version":    versionCommand,
	"delete":     deleteCommand,
	"delete-all": deleteAllCommand,
	"list":       listCommand,
	"slist":      sListCommand,
	"add-var":    addVarCommand,
	"add-cmd":    addCmdCommand,
}

func IsValidCommand(command string) bool {
	_, exists := commands[command]
	return exists
}

func addVarCommand(ctx *GassContext) (message string, err error) {
	if ctx.scenario == "" {
		return "", errors.New("missed scenario name to add variable")
	}

	if len(ctx.params) == 0 {
		return "", errors.New("missed variable name to add")
	}

	//TODO: Check uniqueness of variable
	//TODO: Check if variable name is correct

	ctx.CurrentScenario().Variables = append(ctx.CurrentScenario().Variables, ctx.params[0])
	return fmt.Sprintf("Variable %v added to scenario %v", ctx.params[0], ctx.scenario), nil
}

func addCmdCommand(ctx *GassContext) (message string, err error) {
	if ctx.scenario == "" {
		return "", errors.New("missed scenario name to add command")
	}

	if len(ctx.params) == 0 {
		return "", errors.New("missed command text to add")
	}

	ctx.CurrentScenario().Commands = append(ctx.CurrentScenario().Commands, ctx.params[0])
	return fmt.Sprintf("New command added to scenario %v", ctx.scenario), nil
}

func clearCommand(ctx *GassContext) (message string, err error) {
	if ctx.scenario == "" {
		return "", errors.New("missed scenario name to clear")
	}
	ctx.CurrentScenario().Clear()
	return fmt.Sprintf("Scenario %v is empty now", ctx.scenario), nil
}

func clearAllCommand(ctx *GassContext) (message string, err error) {
	for _, sc := range ctx.lib {
		sc.Clear()
	}
	return "All scenarios are empty now\n", nil
}

func deleteCommand(ctx *GassContext) (message string, err error) {
	if ctx.scenario == "" {
		return "", errors.New("missed scenario name to delete")
	}
	delete(ctx.lib, ctx.scenario)
	return fmt.Sprintf("Scenario %v was deleted", ctx.scenario), nil
}

func deleteAllCommand(ctx *GassContext) (message string, err error) {
	for name := range ctx.lib {
		delete(ctx.lib, name)
	}
	return "All scenarios were deleted\n", nil
}

func listCommand(ctx *GassContext) (message string, err error) {
	if ctx.scenario != "" {
		message += "---------------------------------------------\n"
		message += fmt.Sprintf("Scenario %v: \n%v", ctx.scenario, ctx.CurrentScenario().Text())
	}

	for _, sc := range ctx.lib {
		message += "---------------------------------------------\n"
		message += fmt.Sprintf("Scenario %v: \n%v", sc.Name, sc.Text())
	}
	message += "---------------------------------------------\n"
	return message, nil
}

func sListCommand(ctx *GassContext) (message string, err error) {
	message = "List of Scenarios:\n"
	for name := range ctx.lib {
		message += fmt.Sprintf("   %v \n", name)
	}
	return message, nil
}

func helpCommand(ctx *GassContext) (message string, err error) {
	return HELP_TEXT, nil
}

func versionCommand(ctx *GassContext) (message string, err error) {
	return VERSION_TEXT, nil
}

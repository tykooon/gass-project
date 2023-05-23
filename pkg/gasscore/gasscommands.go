package gasscore

var commands = map[string]func(string, []string) error{
	"clear-all":  nil,
	"clear":      clearCommand,
	"help":       nil,
	"version":    nil,
	"delete":     nil,
	"delete-all": nil,
	"check":      nil,
	"check-all":  nil,
	"info":       nil,
	"list":       nil,
	"add-var":    nil,
	"add-cmd":    nil,
}

func clearCommand(scenario string, params []string) error {
	return nil
}

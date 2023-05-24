package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/tykooon/gass-project/pkg/gasscore"
)

func main() {
	argCount := len(os.Args)

	if argCount == 1 {
		fmt.Print(DISCLAIMER_TEXT)
		return
	}

	logger := log.New(os.Stdout, "gass:", log.Ltime)

	if !checkFile("gass.yaml") {
		logger.Fatal("missing file 'gass.yaml' with gass scenarios...")
	}

	gassData := viper.New()
	gassData.SetConfigName("gass")
	gassData.SetConfigType("yaml")
	gassData.AddConfigPath(".")
	err := gassData.ReadInConfig()
	if err != nil {
		logger.Fatal("Cannot read data from 'gass.yaml'...", err.Error())
	}

	var scLib gasscore.ScenarioLib = readLibrary(gassData)

	Context, err := gasscore.NewContext(os.Args[1:], scLib)

	if err != nil {
		logger.Fatal("Initialization Error ", err.Error())
	}

	res, err := Context.Execute()
	if err != nil {
		logger.Fatal("Execution Error ", err.Error())
	}
	logger.Println(res)

	logger.Println(Context)
}

func checkFile(fname string) bool {
	_, err := os.Stat(fname)
	return err == nil || !os.IsNotExist(err)
}

func readLibrary(data *viper.Viper) (lib map[string]*gasscore.Scenario) {
	lib = make(map[string]*gasscore.Scenario)
	scMap := data.GetStringMap("scenarios")
	for key := range scMap {
		scVariables := data.GetStringSlice("scenarios." + key + ".variables")
		scCommands := data.GetStringSlice("scenarios." + key + ".commands")
		lib[key] = gasscore.NewScenario(key)
		lib[key].AddVariables(scVariables...)
		lib[key].AddCommands(scCommands...)
	}
	return
}

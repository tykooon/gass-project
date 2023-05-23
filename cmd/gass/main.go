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

	logger := log.New(os.Stdout, "gass:", log.Lshortfile)

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

	var scLib gasscore.ScenarioLib = make(map[string]*gasscore.Scenario, 0)
	scMap := gassData.GetStringMap("scenarios")
	for key := range scMap {
		scVariables := gassData.GetStringSlice("scenarios." + key + ".variables")
		scCommands := gassData.GetStringSlice("scenarios." + key + ".commands")
		scen := &gasscore.Scenario{
			Name:      key,
			Variables: scVariables,
			Commands:  scCommands,
		}
		scLib[key] = scen
	}

	Context, err := gasscore.NewContext(os.Args[1:], scLib)

	if err != nil {
		logger.Fatal("Initialization Error ", err.Error())
	}

	fmt.Println(Context)
}

func checkFile(fname string) bool {
	_, err := os.Stat(fname)
	return err == nil || !os.IsNotExist(err)
}

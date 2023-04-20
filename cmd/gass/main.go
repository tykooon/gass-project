package main

import (
	"flag"
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

	flag.Parse()

	if *flagVersion {
		fmt.Printf("Current version: %v\n", CURRENT_VERSION)
	}

	if *flagHelp {
		fmt.Print(HELP_TEXT)
	}

	logger := log.New(os.Stdout, "gass:", log.Lshortfile)

	if err := checkFile("gass.yaml"); err != nil {
		logger.Fatal("missing file 'gass.yaml' with gass scenarios...", err.Error())
	}

	gassData := viper.New()
	gassData.SetConfigName("gass")
	gassData.SetConfigType("yaml")
	gassData.AddConfigPath(".")
	err := gassData.ReadInConfig()
	if err != nil {
		logger.Fatal("missing file 'gass.yaml' with gass scenarios...", err.Error())
	}

	var scLib gasscore.ScenarioLib = make(map[string]gasscore.Scenario, 0)
	scMap := gassData.GetStringMap("scenarios")
	for k := range scMap {
		scVariables := gassData.GetStringSlice("scenarios." + k + ".variables")
		scCommands := gassData.GetStringSlice("scenarios." + k + ".commands")
		scen := gasscore.Scenario{
			Name:      k,
			Variables: scVariables,
			Commands:  scCommands,
		}
		scLib[k] = scen
	}

	fmt.Println(scLib)

}

func checkFile(fname string) error {
	//TODO: check file existence; nil if it exists
	return nil
}

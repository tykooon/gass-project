package main

import "flag"

var ( // Global variables
	flagVersion = flag.Bool("version", false, "get current version of gass")
	flagHelp    = flag.Bool("help", false, "get list of gass commands and options")
)

func init() {

}

const CURRENT_VERSION = "0.0.1"

const HELP_TEXT = `
Usage of gass:
`

const DISCLAIMER_TEXT = `
git assisant (gass) is simple tool for quick execution of short scenarios
with git instuctions. All these scenarios can be found (and also edited)
in the yaml file in the same folder (gass.yaml). These scenarios can be
adjusted and may have additional optiions (e.g. to set commit-messages
or remote-branch names). 

To run scenario just execute gass with the scenario name.
>>gass <scenario-name>

Also, you can view list of options with our help command
>>gass --help (or -h as shortcut)
`

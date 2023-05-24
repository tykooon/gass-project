package main

func init() {

}

const CURRENT_VERSION = "0.0.1"

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

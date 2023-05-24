package gasscore

const VERSION_TEXT = `
Version 0.0.1
`

const HELP_TEXT = `
usage: gass [command?] [scenario?] [parameters...?]

Running scenario from library: 
    scenario_name [...parameters]           

(!Note: if number of parameters less than <scenario> needs - scenario won't proceed)

List of commands:
    add-var <scenario> <name>    add a variable with name <name> to the end of variables list
    add-cmd <scenario> <cmd>     add command with instruction <cmd> to the end commands list
    clear <scenario>             clear variables and commands from <scenario>
    clear-all                    clear all scenarios
    delete <scenario>            delete <scenario> from library 
    delete-all                   delete all scenarios
    help                         show help text
    version                      show version
    list <scenario>              show script of <scenario>. If <scenario> missed - show all scripts
    slist                        short list of scenarios: only names
`

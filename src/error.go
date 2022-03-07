package godo

import (
	"os"
)

const usage = `
usage godo [subcommand]
  Available subcommands
    add <project>	If <project> is omitted, todo will be added to global list
        -t, --title=TITLE		Add a todo with TITLE. If this parameter is not given, launch
					interactive mode
        -d, --description=DESCRIPTION	Add a todo with the given description
        -p, --priority=PRIORITY		Set the priority of the todo (0-9). Defaults to 9 if omitted

    list <project>	If <project> is omitted, all todos will be listed

    remove
`

func FatalWithUsage(message string) {
	os.Stderr.WriteString(message + "\n")
	os.Stdout.WriteString(usage)
	os.Exit(2)
}

func ExitWithUsage() {
	os.Stdout.WriteString(usage)
	os.Exit(0)
}

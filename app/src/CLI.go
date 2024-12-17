package main

import (
	"fmt"
	"regexp"
	"strings"
)

type cliArgs struct {
	seekPath string
	flags    map[string]bool
}

func defaultCLIargs() cliArgs {
	return cliArgs{
		flags: make(map[string]bool),
	}
}

func parseArgs(args []string) cliArgs {

	cliArgs := defaultCLIargs()
	re := regexp.MustCompile(`^--([a-zA-Z0-9-_]+)(=(.*))?$`)

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if matches := re.FindStringSubmatch(arg); matches != nil {
			flagName := matches[1]
			flagValue := matches[3]

			switch flagName {
			case "name":
				if flagValue != "" {
					cliArgs.seekPath = flagValue
				} else if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
					cliArgs.seekPath = args[i+1]
					i++
				} else {
					fmt.Println("Error: --name requires a value")
				}
			default:
				cliArgs.flags[flagName] = true
			}
		}
	}
	return cliArgs
}

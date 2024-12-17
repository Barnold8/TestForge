package main

import (
	"fmt"
	"os"
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

	if cliArgs.seekPath == "" {

		var decision string

		fmt.Println("WARNING: No path was specified. This could lead to unintended results, such as generating tests for directories you did not intend. This is especially risky if you are running this program from the top level of a storage drive.\n\nTry specifying a path with one of the following options:\n--path {your path here} or --path={your path here}")
		fmt.Println("Run anyway? y/n: ")

		fmt.Scanln(&decision)

		decision = strings.ToLower(decision)

		if decision == "y" {
			cliArgs.seekPath = "."
		} else {
			os.Exit(1)
		}

	}

	return cliArgs
}

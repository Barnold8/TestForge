package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type cliArgs struct {
	seekPath   string
	flags      map[string]bool
	ignoreList []string
}

func defaultCLIargs() cliArgs {
	return cliArgs{
		flags: make(map[string]bool),
	}
}

func parseArgs(args []string) cliArgs {

	cliArgs := defaultCLIargs()
	re := regexp.MustCompile(`^--([a-zA-Z0-9-_]+)(=(.*))?$`)
	var ignoreList []string

	for i := 0; i < len(args); i++ {

		arg := args[i]
		arg = strings.ToLower(arg)

		if matches := re.FindStringSubmatch(arg); matches != nil {

			flagName := matches[1]
			flagValue := matches[3]

			switch flagName {
			case "path":
				if flagValue != "" { // case where user uses = in flag
					cliArgs.seekPath = flagValue
				} else if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
					cliArgs.seekPath = args[i+1]
					i++
				} else {
					fmt.Println("Error: --path requires a value")
				}

			case "ignore":

				if flagValue != "" { // case where user uses = in flag
					ignoreList = append(ignoreList, flagValue)

					for i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
						ignoreList = append(ignoreList, args[i+1])
						i++
					}

				} else if !strings.HasPrefix(args[i+1], "-") {

					for i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
						ignoreList = append(ignoreList, args[i+1])
						i++
					}

				} else {
					fmt.Println("Error: --ignore requires a value(s)")
				}

			case "help":

				fmt.Println("TODO: write help information")

			default:
				cliArgs.flags[flagName] = true
			}
		}
	}

	if len(ignoreList) > 0 {
		cliArgs.ignoreList = ignoreList
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

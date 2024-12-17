package main

import (
	"fmt"
	"regexp"
)

func parseArgs(args []string) {
	re := regexp.MustCompile(`^(--\w+(-\w+)*=?|\-\w)(.*)?$`)

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if matches := re.FindStringSubmatch(arg); matches != nil {
			flag := matches[1]

			switch flag {
			case "--allow":
				// allow = true
			case "--verbose":
				// verbose = true
			case "--name":
				// Handle flags with a value (e.g., --name=value or --name value)
				if i+1 < len(args) && args[i+1][0] != '-' {
					// name = args[i+1]
					i++ // Skip the next argument
				} else {
					fmt.Println("Error: --name requires a value")
				}
			default:
				fmt.Printf("Unknown flag: %s\n", flag)
			}
		}
	}
}

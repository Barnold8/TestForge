package main

import (
	"os"
	"strings"
)

func readFileToLines(path string) []string {
	file, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	contentsByNewLine := strings.Split(string(file), "\n")

	return contentsByNewLine
}

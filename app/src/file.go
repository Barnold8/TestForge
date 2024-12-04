package main

import (
	"os"
	"strings"
)

func readFileToLines(path string) ([]string, error) {
	file, err := os.ReadFile(path)

	if err != nil {
		return []string{""}, err
	}

	contentsByNewLine := strings.Split(string(file), "\n")

	return contentsByNewLine, err
}

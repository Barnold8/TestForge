package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var FILEPERM fs.FileMode = 0644

func readFileToLines(path string) ([]string, error) {
	file, err := os.ReadFile(path)

	if err != nil {
		return []string{""}, err
	}

	contentsByNewLine := strings.Split(string(file), "\n")

	return contentsByNewLine, err
}

func seekGoFiles(path string, overWrite bool) ([]string, error) {

	var goFiles []string
	fileMap := make(map[string]string)

	err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {

		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		if d == nil {
			return nil
		}

		if !d.IsDir() && strings.HasSuffix(d.Name(), ".go") {
			path = strings.ReplaceAll(path, "\\", "/")
			if strings.HasSuffix(d.Name(), "_test.go") && !overWrite { // if overwrite is false, remove all .go files that have a _test counterpart
				delete(fileMap, strings.Split(path, "_test")[0]+".go") // thankfully this function is written well so no error avoidance is needed
			} else if !strings.HasSuffix(d.Name(), "_test.go") {

				fileMap[path] = path
			}

		}

		return nil
	})

	for key := range fileMap {
		goFiles = append(goFiles, key)
	}

	return goFiles, err
}

func writeFile(path string, contents string) error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, FILEPERM)
	if err != nil {
		file.Close()
		return err
	}
	_, err = file.WriteString(contents)

	if err != nil {
		file.Close()
		return err
	}

	file.Close()

	return nil
}

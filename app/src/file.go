package main

import (
	"os"
	"path/filepath"
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

func seekGoFiles(path string, overWrite bool) ([]string, error) {

	var goFiles []string
	fileMap := make(map[string]string)

	err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".go") {

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

// TODO:

// Have a directory that we look into, this will probably come in through a CLI arg
// 	Walk through that directory, even into sub directories and log all the go files that dont have a test version of themselves (maybe add a override to overwrite existing test files)
//

// Make a test file for a given .go file
// 	go through file strings
// 		if func in string
// 			make a string[] to store all strings up until we find a "{"
// 				then build that string array to a singular string for easier processing
// 					get function args from all stuff after "(" and before ")", this can then be split up via "," as a delimeter
//					get function returns from all stuff after "(" and before "{"
//

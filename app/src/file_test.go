package main

import (
	"fmt"
	"io/fs"
	"os"
	"testing"
)

var FILEPERM fs.FileMode = 0600

func directoryExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

func writeTestFile(path string, fileContents string) error {

	// Attempt to generate TestFiles
	err := os.WriteFile(path, []byte(fileContents), FILEPERM)
	// Attempt to generate TestFiles

	return err
}

func writeTestDirectory(path string) error {

	exists, err := directoryExists(path)

	if err != nil {
		return fmt.Errorf("Error - writeTestDirectory: An error was caught in directoryExists %s\n", err)
	}
	if exists == true {
		return fmt.Errorf("Error - writeTestDirectory: Tried writing test directory (%s), but it already exists\n", path)
	} else {
		possibleError := os.Mkdir(path, FILEPERM)
		if possibleError != nil {
			return possibleError
		}
	}

	return nil
}

func cleanup(path string) bool {

	return false
}

func TestReadFileToLines(t *testing.T) {

	// pre testing error catching
	var errors []error
	errors = append(errors, writeTestDirectory("test"))
	errors = append(errors, writeTestFile("test/main.go", "package main\n import \"fmt\" func main(){\n\n\tfmt.Println(\"Hello world!\")\n\n}")) // Make the temp files

	for _, err := range errors {

		if err != nil {
			t.Errorf("Error - TestReadFileToLines: %s", err)
		}

	}
	// pre testing error catching

	// file, err := os.ReadFile()

	// if err != nil {
	// 	panic(err)
	// }

	// contentsByNewLine := strings.Split(string(file), "\n")

	// return contentsByNewLine
}

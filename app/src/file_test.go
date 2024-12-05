package main

import (
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"strings"
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
	err := os.WriteFile(path, []byte(fileContents), FILEPERM)
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

func cleanup(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func TestReadFileToLines(t *testing.T) {

	PATH := "tests/"

	// pre testing error catching
	var errors []error
	errors = append(errors, writeTestDirectory("tests"))
	errors = append(errors, writeTestFile(PATH+"main.go", "package main\nimport \"fmt\"\n\n func main(){\n\n\tfmt.Println(\"Hello world!\")\n\n}")) // Make the temp files
	errors = append(errors, writeTestFile(PATH+"main.txt", "Hello world\n\nThis is my test file!"))
	errors = append(errors, writeTestFile(PATH+"empty.txt", ""))

	for _, err := range errors {

		if err != nil {
			cleanup(PATH)
			t.Errorf("Error - TestReadFileToLines: %s", err)
		}

	}
	// pre testing error catching

	tests := []struct {
		name     string
		path     string
		expected []string
	}{
		{"Read main.go", "tests/main.go", strings.Split("package main\nimport \"fmt\"\n\n func main(){\n\n\tfmt.Println(\"Hello world!\")\n\n}", "\n")},
		{"Read main.txt", "tests/main.txt", strings.Split("Hello world\n\nThis is my test file!", "\n")},
		{"Read empty.txt", "tests/empty.txt", strings.Split("", "\n")},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result, err := readFileToLines(tc.path)

			if err != nil {
				cleanup(PATH)
				t.Errorf("Error - TestReadFileToLines: %s", err)
			} else if reflect.DeepEqual(result, tc.expected) == false {
				cleanup(PATH)
				t.Error("Error - TestReadFileToLines: result and tc.expected do not match")
				t.Errorf("== RESULT ==\n\n%s\n\n == EXPECTED ==\n\n%s\n\n", result, tc.expected)
			}

		})
	}

	// cleanup after testing

	cleanup(PATH)

	// cleanup after testing

}

func TestSeekGoFiles(t *testing.T) {
	PATH := "tests/"

	// pre testing error catching
	var errors []error
	errors = append(errors, writeTestDirectory("tests"))
	errors = append(errors, writeTestDirectory("tests/test1"))
	errors = append(errors, writeTestDirectory("tests/test2"))
	errors = append(errors, writeTestDirectory("tests/test3"))
	errors = append(errors, writeTestDirectory("tests/test4"))
	errors = append(errors, writeTestFile(PATH+"main.go", ""))
	errors = append(errors, writeTestFile(PATH+"main.txt", ""))
	errors = append(errors, writeTestFile(PATH+"empty.txt", ""))

	for _, err := range errors {

		if err != nil {
			t.Errorf("Error - TestReadFileToLines: %s", err)
		}

	}
	// pre testing error catching

	// cleanup after testing

	cleanup(PATH)

	// cleanup after testing

}

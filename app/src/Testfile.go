package main

import (
	"os"
	"testing"
)

func writeTestFiles() bool {

	var errors []error

	errors = append(errors, os.WriteFile("test.txt", []byte("hello world"), 0644))
	errors = append(errors, os.WriteFile("main.go", []byte("package main\n\nfunc main(){\n fmt.Println(\" Hello World \")"), 0644))

	for _, err := range errors {

		if err != nil {
			return false
		}
	}
	return true
}

func cleanTestFiles() bool {

	return false
}

func TestreadFileToLines(t *testing.T) {

	writeTestFiles() // Make the temp files

	// file, err := os.ReadFile()

	// if err != nil {
	// 	panic(err)
	// }

	// contentsByNewLine := strings.Split(string(file), "\n")

	// return contentsByNewLine
}

package main

import (
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	PATH := "tests/"
	// pre testing error catching
	var errors []error
	errors = append(errors, writeTestDirectory("tests"))
	errors = append(errors, writeTestFile(PATH+"main.go", "package main\nimport \"fmt\"\n\n func main(){\n\n\tfmt.Println(\"Hello world!\")\n\n}")) // Make the temp files
	errors = append(errors, writeTestFile(PATH+"main.txt", "Hello world\n\nThis is my test file!"))
	errors = append(errors, writeTestFile(PATH+"empty.txt", ""))
	errors = append(errors, writeTestFile(PATH+"test1.go", "package main\nimport \"fmt\"\n\n func printNumbers(int x, int y){\n\n\tfmt.Printf(\"%d %d\",x,y))")) // Make the temp files

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
		expected goFile
	}{
		{"main.go", "tests/main.go", goFile{
			"tests/main.go",
			"package main",
			[]goFunction{
				goFunction{
					"main",
					nil,
					"",
				},
			},
		}},
		{"main.txt", "tests/main.txt", goFile{
			"tests/main.txt",
			"",
			nil,
		}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := ParseFile(tc.path)

			if !reflect.DeepEqual(tc.expected, result) {
				cleanup(PATH)

				t.Errorf("Error - TestParseFile\n\n")

				if result.filePath != tc.expected.filePath {
					t.Errorf("Error - TestParseFile - Mismatched filePaths\n\n %s (result) VS %s (expected)\n\n", result.filePath, tc.expected.filePath)
				}

				if result.filepackage != tc.expected.filepackage {
					t.Errorf("Error - TestParseFile - Mismatched filePackages\n\n %s (result) VS %s (expected)\n\n", result.filePath, tc.expected.filePath)
				}

				if !reflect.DeepEqual(tc.expected.fileFunctions, result.fileFunctions) {

					t.Errorf("Error - TestParseFile - Mismatched function Arrays")

					if len(tc.expected.fileFunctions) != len(result.fileFunctions) {
						t.Errorf("Error - TestParseFile - Mismatched function Array lengths\n\n %d (result) VS %d (expected)\n\n", len(result.fileFunctions), len(tc.expected.fileFunctions))
						return
					}
					t.Errorf("Error - TestParseFile - Mismatched function array contents")
					for i := range result.fileFunctions {

						t.Errorf("\n\n====INDEX %d===\n", i)

						if result.fileFunctions[i].funcName != tc.expected.fileFunctions[i].funcName {
							t.Errorf("Function name %s (result) VS Function name %s (expected)\n\n", result.fileFunctions[i].funcName, tc.expected.fileFunctions[i].funcName)
						}

						if result.fileFunctions[i].funcReturn != tc.expected.fileFunctions[i].funcReturn {
							t.Errorf("Function return %s (result) VS Function return %s (expected)\n\n", result.fileFunctions[i].funcName, tc.expected.fileFunctions[i].funcName)
						}

						if !reflect.DeepEqual(result.fileFunctions[i].funcArgs, tc.expected.fileFunctions[i].funcArgs) {
							if len(result.fileFunctions[i].funcArgs) != len(tc.expected.fileFunctions[i].funcArgs) {
								t.Errorf("Error - TestParseFile - Mismatched function arg lengths\n\n %d (result) VS %d (expected)\n\n", len(result.fileFunctions[i].funcArgs), len(tc.expected.fileFunctions[i].funcArgs))
							}
						}

					}

					t.Errorf("Error - TestParseFile - Printing raw output\n\nRESULT: LENGTH[%d] DATA %v IS_NIL %v\n\nEXPECTED: LENGTH[%d] DATA %v IS_NIL %v\n\n", len(result.fileFunctions), result.fileFunctions, result.fileFunctions == nil, len(tc.expected.fileFunctions), tc.expected.fileFunctions, tc.expected.fileFunctions == nil)

				}

			}

		})
	}

	cleanup(PATH)

}

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
	errors = append(errors, writeTestFile(PATH+"main.go", "package main\nimport \"fmt\"\n\nfunc main(){\n\n\tfmt.Println(\"Hello world!\")\n\n}")) // Make the temp files
	errors = append(errors, writeTestFile(PATH+"main.txt", "Hello world\n\nThis is my test file!"))
	errors = append(errors, writeTestFile(PATH+"empty.txt", ""))
	errors = append(errors, writeTestFile(PATH+"test1.go", "package tests\nimport \"fmt\"\n\nfunc printNumbers(x int,y int){\n\n\tfmt.Printf(\"%d %d\",x,y)\n}")) // Make the temp files
	errors = append(errors, writeTestFile(PATH+"test2.go", "package testingTHIS\nimport \"fmt\"\n\nfunc a_REALLY_stupidFunction_NaMe(fff []error) int {\n\n\tfmt.Printf(\"%d %d\",x,y)\n}"))

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
		{"empty.txt", "tests/empty.txt", goFile{
			"tests/empty.txt",
			"",
			nil,
		}},
		{"test1.go", "tests/test1.go", goFile{
			"tests/test1.go",
			"package tests",
			[]goFunction{
				goFunction{
					"printNumbers",
					[]string{"x int", "y int"},
					"",
				},
			},
		}},
		{"test2.go", "tests/test2.go", goFile{
			"tests/test2.go",
			"package testingTHIS",
			[]goFunction{
				goFunction{
					"a_REALLY_stupidFunction_NaMe",
					[]string{"fff []error"},
					"int",
				},
			},
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
							t.Errorf("Function return %s (result) VS Function return %s (expected)\n\n", result.fileFunctions[i].funcReturn, tc.expected.fileFunctions[i].funcReturn)
						}

						if !reflect.DeepEqual(result.fileFunctions[i].funcArgs, tc.expected.fileFunctions[i].funcArgs) {
							if len(result.fileFunctions[i].funcArgs) != len(tc.expected.fileFunctions[i].funcArgs) {
								t.Errorf("Error - TestParseFile - Mismatched function arg lengths\n\n %d (result) VS %d (expected)\n\n", len(result.fileFunctions[i].funcArgs), len(tc.expected.fileFunctions[i].funcArgs))
								return
							}

							for j := range result.fileFunctions[i].funcArgs {

								t.Errorf("Result arg[%d]: %s VS EXPECTED arg[%d]: %s", j, result.fileFunctions[i].funcArgs[j], j, tc.expected.fileFunctions[i].funcArgs[j])
							}

						}
					}
				}

			}

		})
	}

	cleanup(PATH)

}

func TestParseFunction(t *testing.T) {

	test1 := []string{`func TestParseFunction(t *testing.T) {
		tests := []struct {
			name     string
			path     string
			expected goFile
		}{}
	
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
	
			})
		}
	}`}

	test2 := []string{`func addition(x int,y int) int {
		return x + y
	}`}

	test3 := []string{``}

	test4 := []string{`funcfuncfunc(x int,y int) int {
		return x + y
	}`}
	test5 := []string{`def pythonFunc(x,y,z):
		return x + y * z
	`}

	test6 := []string{`func _SomEThing_CoMpleX(ASILLYCOMPLEXDATATYPE datatypecomplex,a complexhandler) (error,complex,int,float) {
		return x + y
	}`}

	tests := []struct {
		name           string
		functionString *[]string
		expected       goFunction
	}{

		{"Test 1", &test1, goFunction{"TestParseFunction", []string{"t *testing.T"}, ""}},
		{"Test 2", &test2, goFunction{"addition", []string{"x int", "y int"}, "int"}},
		{"Test 3", &test3, goFunction{}},
		{"Test 4", &test4, goFunction{}},
		{"Test 5", &test5, goFunction{}},
		{"Test 6", &test6, goFunction{"_SomEThing_CoMpleX", []string{"ASILLYCOMPLEXDATATYPE datatypecomplex", "a complexhandler"}, "(error,complex,int,float)"}},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := parseFunction(tc.functionString, 0)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Error - TestParseFunction Resulting goFunction does not match the expected goFunction. See details below:\n\nRESULT:  %v\nEXPECTED:%v  ", result, tc.expected)
			}

		})
	}
}

func TestRemoveEmptyFuncs(t *testing.T) {

	tests := []struct {
		name              string
		unParsedFunctions []goFunction
		expected          []goFunction
	}{

		{"Test 1",
			[]goFunction{
				goFunction{"addition", []string{"x int", "y int"}, "int"},
			}, []goFunction{
				goFunction{"addition", []string{"x int", "y int"}, "int"}},
		},
		{"Test 2", []goFunction{}, []goFunction{}},
		{"Test 1",
			[]goFunction{
				goFunction{"editInPlace", []string{}, ""},
				goFunction{},
				goFunction{"editOutOfPlace", []string{}, ""},
				goFunction{},
				goFunction{},
				goFunction{"abcxyz", []string{"1", "2", "3"}, "error"},
				goFunction{},
				goFunction{},
			}, []goFunction{
				goFunction{"editInPlace", []string{}, ""},
				goFunction{"editOutOfPlace", []string{}, ""},
				goFunction{"abcxyz", []string{"1", "2", "3"}, "error"},
			},
		},
		{"Test 4", []goFunction{goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}, goFunction{}}, []goFunction{}},
		{"Test 5", []goFunction{goFunction{"main", []string{}, ""}}, []goFunction{goFunction{"main", []string{}, ""}}},
		{"Test 6", []goFunction{
			goFunction{"Abcdefghi", []string{"_ccc_zz_A_DFf_g_gh_zxv", "aa", "vvbbbb", ""}, ""},
			goFunction{"", []string{"", ""}, ""},
			goFunction{"", []string{"", ""}, ""},
			goFunction{"", []string{"", ""}, ""},
			goFunction{},
			goFunction{},
			goFunction{},
			goFunction{},
			goFunction{},
		},
			[]goFunction{goFunction{"Abcdefghi", []string{"_ccc_zz_A_DFf_g_gh_zxv", "aa", "vvbbbb", ""}, ""},
				goFunction{"", []string{"", ""}, ""},
				goFunction{"", []string{"", ""}, ""},
				goFunction{"", []string{"", ""}, ""}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			removeEmptyFuncs(&tc.unParsedFunctions)

			if !reflect.DeepEqual(tc.unParsedFunctions, tc.expected) {
				t.Errorf("Error - TestRemoveEmptyFuncs: Result does not match expected.\nRESULT:  %v\nEXPECTED: %v\n", tc.unParsedFunctions, tc.expected)
			}
		})
	}
}

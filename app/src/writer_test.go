package main

import (
	"testing"
)

func TestWriteTests(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test 1", "input", "output"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// some way to get a result
			result := "this is an example of a result"
			if result != tc.expected {
				t.Errorf("This is an example of an error!")
			}
		})
	}
}

func TestFunctionArgsToString(t *testing.T) {

	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{"Test 1", []string{}, ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := functionArgsToString(&tc.input)

			if result != tc.expected {
				t.Errorf("ERROR")
			}

		})
	}
}

func TestGoFunctionsToString(t *testing.T) {

	cases := "\n\n\ttests := []struct {\n\t\tname string\n\t\tinput string\n\t\texpected string\n\t}{\n\t\t{\"Test 1\",\"input\",\"output\"},\n\t}\n\tfor _, tc := range tests {\n\t\tt.Run(tc.name, func(t *testing.T){\n\t\t\t// some way to get a result\n\t\t\tresult := \"this is an example of a result\"\n\t\t\tif result != tc.expected{\n\t\t\t\t t.Errorf(\"This is an example of an error!\")\n\t\t\t}\n\t\t})\n\t}\n}"

	tests := []struct {
		name     string
		input    []goFunction
		cliArgs  cliArgs
		expected string
	}{
		{"Test 1", []goFunction{}, defaultCLIargs(), ""},
		{"Test 2", []goFunction{
			goFunction{"addition", []string{"x int", "y int"}, "int"},
		},
			cliArgs{flags: map[string]bool{
				"cases": true,
			}},
			"\n\nfunc TestAddition(t *testing.T) {" + cases,
		},
		{"Test 3", []goFunction{
			goFunction{"addition", []string{"x int", "y int"}, "int"},
			goFunction{"minus", []string{"x int", "y int"}, "int"},
		},
			cliArgs{flags: map[string]bool{
				"cases": true,
			}},
			"\n\nfunc TestAddition(t *testing.T) {" + cases + "\n\nfunc TestMinus(t *testing.T) {" + cases,
		},
		{"Test 4", []goFunction{
			goFunction{"someWhackyFunctioniDontKnow", []string{"throwYouOff error", "somerandompseudotype ptype"}, "ptype"},
			goFunction{"literallyNothing", []string{}, ""},
		},
			cliArgs{flags: map[string]bool{
				"cases": true,
			}},
			"\n\nfunc TestSomeWhackyFunctioniDontKnow(t *testing.T) {" + cases + "\n\nfunc TestLiterallyNothing(t *testing.T) {" + cases,
		},
		{"Test 5", []goFunction{}, defaultCLIargs(), ""},
		{"Test 6", []goFunction{
			goFunction{"addition", []string{"x int", "y int"}, "int"},
		},
			cliArgs{flags: map[string]bool{
				"cases": false,
			}},
			"\n\nfunc TestAddition(t *testing.T) {\n\n}",
		},
		{"Test 7", []goFunction{
			goFunction{"addition", []string{"x int", "y int"}, "int"},
			goFunction{"minus", []string{"x int", "y int"}, "int"},
		},
			cliArgs{flags: map[string]bool{
				"cases": false,
			}},
			"\n\nfunc TestAddition(t *testing.T) {\n\n}" + "\n\nfunc TestMinus(t *testing.T) {\n\n}",
		},
		{"Test 8", []goFunction{
			goFunction{"someWhackyFunctioniDontKnow", []string{"throwYouOff error", "somerandompseudotype ptype"}, "ptype"},
			goFunction{"literallyNothing", []string{}, ""},
		},
			cliArgs{flags: map[string]bool{
				"cases": false,
			}},
			"\n\nfunc TestSomeWhackyFunctioniDontKnow(t *testing.T) {\n\n}" + "\n\nfunc TestLiterallyNothing(t *testing.T) {\n\n}",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// some way to get a result

			result := goFunctionsToString(&tc.input, tc.cliArgs)

			if result != tc.expected {
				t.Errorf("RESULT\n\n %s\n\nEXPECTED\n\n %s", result, tc.expected)
			}

		})
	}
}

func TestCapitalizeFunctionName(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test 1", "input", "output"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// some way to get a result
			result := "this is an example of a result"
			if result != tc.expected {
				t.Errorf("This is an example of an error!")
			}
		})
	}
}

func TestFormatFileName(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test 1", "input", "output"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// some way to get a result
			result := "this is an example of a result"
			if result != tc.expected {
				t.Errorf("This is an example of an error!")
			}
		})
	}
}

func TestGatherFiles(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test 1", "input", "output"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// some way to get a result
			result := "this is an example of a result"
			if result != tc.expected {
				t.Errorf("This is an example of an error!")
			}
		})
	}
}

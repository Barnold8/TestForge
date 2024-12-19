package main

import (
	"reflect"
	"testing"
)

func TestDefaultCLIargs(t *testing.T) {

	tests := []struct {
		name     string
		expected cliArgs
	}{
		{"Test 1", cliArgs{
			flags: make(map[string]bool),
		}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// some way to get a result
			result := defaultCLIargs()

			if !(result.flags != nil && tc.expected.flags != nil &&
				result.ignoreList == nil && tc.expected.ignoreList == nil &&
				result.seekPath == "" && tc.expected.seekPath == "") {
				t.Errorf("Error - TestDefaultCLIargs\n\nResult was %v when %v was expected\n\n", result, tc.expected)
			} else {
				if len(result.flags) != len(tc.expected.flags) {
					t.Errorf("Error - TestDefaultCLIargs\n\nResult was %v when %v was expected\n\n", result, tc.expected)
				} else if len(result.flags) > 0 {
					t.Errorf("Error - TestDefaultCLIargs\n\nAn underlying issue with the defaultCLIArgs function was found, the length of this is expected to be 0 when %d was found\n\n", len(tc.expected.flags))
				}
			}

		})
	}
}

func TestParseArgs(t *testing.T) {

	tests := []struct {
		name     string
		input    []string
		expected cliArgs
	}{

		{"Test 1", []string{"--path", "test", "--overwrite", "--ignore", "C:\\hello\\world"}, cliArgs{
			seekPath: "test",
			flags: map[string]bool{
				"overwrite": true,
			},
			ignoreList: []string{"C:\\hello\\world"},
		}},
		{"Test 2", []string{"--path", "something\\a\\bit\\far\\away", "--overwrite", "--someflag", "--a", "--b"}, cliArgs{
			seekPath: "something\\a\\bit\\far\\away",
			flags: map[string]bool{
				"overwrite": true,
				"someflag":  true,
				"a":         true,
				"b":         true,
			},
			ignoreList: nil,
		}},

		{"Test 3", []string{"--path=another/test/path", "--overwrite", "--ignore=file1.txt file2.txt"}, cliArgs{
			seekPath: "another/test/path",
			flags: map[string]bool{
				"overwrite": true,
			},
			ignoreList: []string{"file1.txt file2.txt"},
		}},

		{"Test 4", []string{"--path", "test", "--ignore", "ignore1", "ignore2", "ignore3"}, cliArgs{
			seekPath: "test",
			flags:    map[string]bool{},
			ignoreList: []string{
				"ignore1",
				"ignore2",
				"ignore3",
			},
		}},

		{"Test 5", []string{"--Path", "upper/Case", "--Overwrite", "--SomeFlag"}, cliArgs{
			seekPath: "upper/Case",
			flags: map[string]bool{
				"overwrite": true,
				"someflag":  true,
			},
			ignoreList: nil,
		}},

		{"Test 6", []string{"--path", "valid/path", "--overwrite", "--someflag"}, cliArgs{
			seekPath: "valid/path",
			flags: map[string]bool{
				"overwrite": true,
				"someflag":  true,
			},
			ignoreList: nil,
		}},

		{"Test 7", []string{"--path", "valid/path", "--someflag", "--anotherflag"}, cliArgs{
			seekPath: "valid/path",
			flags: map[string]bool{
				"someflag":    true,
				"anotherflag": true,
			},
			ignoreList: nil,
		}},

		{"Test 8", []string{"--path", "valid/path", "--ignore", "file1.txt", "--invalidflag"}, cliArgs{
			seekPath: "valid/path",
			flags: map[string]bool{
				"invalidflag": true,
			},
			ignoreList: []string{"file1.txt"},
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := parseArgs(tc.input)

			if !(result.seekPath == tc.expected.seekPath && reflect.DeepEqual(result.flags, tc.expected.flags) && reflect.DeepEqual(result.ignoreList, tc.expected.ignoreList)) {
				t.Errorf("\n\nError - TestDefaultCLIargs\n\n")

				if result.seekPath != tc.expected.seekPath {
					t.Errorf("\n\nResult seekpath is %swhen %s was expected\n\n", result.seekPath, tc.expected.seekPath)
				}

				if !reflect.DeepEqual(result.flags, tc.expected.flags) {
					t.Errorf("\n\nResult flags are \n\n%v\n\nwhen \n\n%v was expected\n\n", result.flags, tc.expected.flags)
				}

				if !reflect.DeepEqual(result.ignoreList, tc.expected.ignoreList) {
					t.Errorf("\n\nResult ignoreList is \n\n%v\n\nwhen \n\n%v was expected\n\n", result.ignoreList, tc.expected.ignoreList)
				}
			}

		})
	}
}

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
				t.Errorf("Error - TestGoFunctionsToString: \n\nRESULT\n\n %s\n\nEXPECTED\n\n %s", result, tc.expected)
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
		{"Test 1", "someFunction", "SomeFunction"},
		{"Test 2", "QuiteliteralOtherfunction", "QuiteliteralOtherfunction"},
		{"Test 3", "anotherExampleFunction", "AnotherExampleFunction"},
		{"Test 4", "thisIsATest", "ThisIsATest"},
		{"Test 5", "parseURL", "ParseURL"},
		{"Test 6", "getHTTPRequest", "GetHTTPRequest"},
		{"Test 7", "generateID", "GenerateID"},
		{"Test 8", "toLowercase", "ToLowercase"},
		{"Test 9", "convertStringToInt", "ConvertStringToInt"},
		{"Test 10", "processFile", "ProcessFile"},
		{"Test 11", "findAndReplace", "FindAndReplace"},
		{"Test 12", "sortArray", "SortArray"},
		{"Test 13", "filterData", "FilterData"},
		{"Test 14", "validateUserInput", "ValidateUserInput"},
		{"Test 15", "handleError", "HandleError"},
		{"Test 16", "readConfigFile", "ReadConfigFile"},
		{"Test 17", "saveToDatabase", "SaveToDatabase"},
		{"Test 18", "deleteRecord", "DeleteRecord"},
		{"Test 19", "fetchData", "FetchData"},
		{"Test 20", "calculateSum", "CalculateSum"},
		{"Test 21", "openConnection", "OpenConnection"},
		{"Test 22", "closeConnection", "CloseConnection"},
		{"Test 23", "parseJSON", "ParseJSON"},
		{"Test 24", "convertToCamelCase", "ConvertToCamelCase"},
		{"Test 25", "isPrimeNumber", "IsPrimeNumber"},
		{"Test 26", "checkPalindrome", "CheckPalindrome"},
		{"Test 27", "reverseString", "ReverseString"},
		{"Test 28", "splitSentence", "SplitSentence"},
		{"Test 29", "mergeArrays", "MergeArrays"},
		{"Test 30", "capitalizeWords", "CapitalizeWords"},
		{"Test 31", "encryptPassword", "EncryptPassword"},
		{"Test 32", "decodeBase64", "DecodeBase64"},
		{"Test 33", "generateToken", "GenerateToken"},
		{"Test 34", "validateEmail", "ValidateEmail"},
		{"Test 35", "processTransaction", "ProcessTransaction"},
		{"Test 36", "findMaximumValue", "FindMaximumValue"},
		{"Test 37", "isEvenNumber", "IsEvenNumber"},
		{"Test 38", "createNewUser", "CreateNewUser"},
		{"Test 39", "sendHTTPRequest", "SendHTTPRequest"},
		{"Test 40", "parseCommandLine", "ParseCommandLine"},
		{"Test 41", "logToConsole", "LogToConsole"},
		{"Test 42", "convertToSnakeCase", "ConvertToSnakeCase"},
		{"Test 43", "formatDate", "FormatDate"},
		{"Test 44", "calculateFactorial", "CalculateFactorial"},
		{"Test 45", "findDuplicateEntries", "FindDuplicateEntries"},
		{"Test 46", "sortByColumn", "SortByColumn"},
		{"Test 47", "getFileExtension", "GetFileExtension"},
		{"Test 48", "resizeImage", "ResizeImage"},
		{"Test 49", "uploadFile", "UploadFile"},
		{"Test 50", "downloadFile", "DownloadFile"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// some way to get a result
			result := capitalizeFunctionName(tc.input)
			if result != tc.expected {
				t.Errorf("Error - TestCapitalizeFunctionName: \n\nRESULT\n\n %s\n\nEXPECTED\n\n %s", result, tc.expected)
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
		{"Test 1", "somefile.go", "somefile_test.go"},
		{"Test 2", "main.go", "main_test.go"},
		{"Test 3", "utils.go", "utils_test.go"},
		{"Test 4", "handlers.go", "handlers_test.go"},
		{"Test 5", "config.go", "config_test.go"},
		{"Test 6", "router.go", "router_test.go"},
		{"Test 7", "middleware.go", "middleware_test.go"},
		{"Test 8", "database.go", "database_test.go"},
		{"Test 9", "models.go", "models_test.go"},
		{"Test 10", "auth.go", "auth_test.go"},
		{"Test 11", "cache.go", "cache_test.go"},
		{"Test 12", "logger.go", "logger_test.go"},
		{"Test 13", "session.go", "session_test.go"},
		{"Test 14", "validator.go", "validator_test.go"},
		{"Test 15", "parser.go", "parser_test.go"},
		{"Test 16", "scheduler.go", "scheduler_test.go"},
		{"Test 17", "worker.go", "worker_test.go"},
		{"Test 18", "notifier.go", "notifier_test.go"},
		{"Test 19", "emailer.go", "emailer_test.go"},
		{"Test 20", "uploader.go", "uploader_test.go"},
		{"Test 21", "downloader.go", "downloader_test.go"},
		{"Test 22", "converter.go", "converter_test.go"},
		{"Test 23", "encryptor.go", "encryptor_test.go"},
		{"Test 24", "decryptor.go", "decryptor_test.go"},
		{"Test 25", "hasher.go", "hasher_test.go"},
		{"Test 26", "comparator.go", "comparator_test.go"},
		{"Test 27", "merger.go", "merger_test.go"},
		{"Test 28", "splitter.go", "splitter_test.go"},
		{"Test 29", "resolver.go", "resolver_test.go"},
		{"Test 30", "matcher.go", "matcher_test.go"},
		{"Test 31", "generator.go", "generator_test.go"},
		{"Test 32", "processor.go", "processor_test.go"},
		{"Test 33", "cleaner.go", "cleaner_test.go"},
		{"Test 34", "formatter.go", "formatter_test.go"},
		{"Test 35", "renderer.go", "renderer_test.go"},
		{"Test 36", "exporter.go", "exporter_test.go"},
		{"Test 37", "importer.go", "importer_test.go"},
		{"Test 38", "analyzer.go", "analyzer_test.go"},
		{"Test 39", "optimizer.go", "optimizer_test.go"},
		{"Test 40", "executor.go", "executor_test.go"},
		{"Test 41", "scheduler.go", "scheduler_test.go"},
		{"Test 42", "controller.go", "controller_test.go"},
		{"Test 43", "dispatcher.go", "dispatcher_test.go"},
		{"Test 44", "observer.go", "observer_test.go"},
		{"Test 45", "monitor.go", "monitor_test.go"},
		{"Test 46", "tracker.go", "tracker_test.go"},
		{"Test 47", "seeder.go", "seeder_test.go"},
		{"Test 48", "crawler.go", "crawler_test.go"},
		{"Test 49", "indexer.go", "indexer_test.go"},
		{"Test 50", "archiver.go", "archiver_test.go"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := formatFileName(tc.input)

			if result != tc.expected {

				t.Errorf("Error - TestGoFunctionsToString: \n\nRESULT\n\n %s\n\nEXPECTED\n\n %s", result, tc.expected)

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

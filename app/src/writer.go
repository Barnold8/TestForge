package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func writeTests(path string, overwrite bool) {

	files := gatherFiles(path, overwrite)

	for _, value := range files {
		var builder strings.Builder

		builder.WriteString(value.filepackage + "\n\nimport(\"testing\")")

		builder.WriteString(goFunctionsToString(&value.fileFunctions, true))

		writeFile(formatFileName(value.filePath), builder.String())

	}

}

func functionArgsToString(args *[]string) string {

	var builder strings.Builder

	for i, value := range *args {

		if i != len(*args)-1 {
			builder.WriteString(fmt.Sprintf("%s,", value))
		} else {
			builder.WriteString(value)
		}

	}

	return builder.String()
}

func goFunctionsToString(function *[]goFunction, writeTestCases bool) string {

	var builder strings.Builder

	for _, value := range *function {

		if writeTestCases {

			builder.WriteString(fmt.Sprintf(
				"\n\nfunc Test%s(t *testing.T) {\n\n\ttests := []struct {\n\t\tname string\n\t\tinput string\n\t\texpected string\n\t}{\n\t\t{\"Test 1\",\"input\",\"output\"},\n\t}\n\tfor _, tc := range tests {\n\t\tt.Run(tc.name, func(t *testing.T){\n\t\t\t// some way to get a result\n\t\t\tresult := \"this is an example of a result\"\n\t\t\tif result != tc.expected{\n\t\t\t\t t.Errorf(\"This is an example of an error!\")\n\t\t\t}\n\t\t})\n\t}\n}",
				capitalizeFunctionName(value.funcName),
			))
		} else {
			builder.WriteString(fmt.Sprintf(
				"\n\nfunc Test%s(t *testing.T) {\n\n}",
				value.funcName,
			))

		}

	}

	return builder.String()
}

func capitalizeFunctionName(name string) string {

	re := regexp.MustCompile(`([a-z])([A-Z])|(^[a-z])`)

	capitalized := re.ReplaceAllStringFunc(name, func(match string) string {
		if len(match) == 1 {

			return strings.ToUpper(match)
		}

		return match[:1] + strings.ToUpper(match[1:])
	})

	r := []rune(capitalized)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func formatFileName(path string) string {
	return strings.Split(path, ".")[0] + "_test.go"
}

func gatherFiles(path string, overwrite bool) []goFile {

	filesPaths, err := seekGoFiles(path, overwrite)

	var goFiles []goFile

	if err != nil {

		fmt.Printf("Error: %s\n", err)
		return []goFile{}

	} else {

		for _, filePath := range filesPaths {
			goFiles = append(goFiles, ParseFile(filePath))
		}

	}
	return goFiles
}

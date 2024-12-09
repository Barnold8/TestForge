package main

import (
	"fmt"
	"regexp"
	"strings"
)

type goFunction struct {
	funcName   string
	funcArgs   []string
	funcReturn string
}

type goFile struct {
	filePath      string
	filepackage   string
	fileFunctions []goFunction
}

func ParseFile(path string) goFile {

	var file = goFile{}
	// var goFunction goFunction
	file.filePath = path
	fileContents, err := readFileToLines(path)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return goFile{}
	}

	for i := 0; i < len(fileContents); i++ {

		if !strings.Contains(fileContents[i], "//") { // doesnt account for strings that have "//" in them, probably wont occur within a function or package declaration

			if strings.Contains(fileContents[i], "package") {
				file.filepackage = fileContents[i]
			}

			if strings.Contains(fileContents[i], "func") {
				file.fileFunctions = append(file.fileFunctions, parseFunction(&fileContents, i))
			}

		}
	}

	return file

}

func parseFunction(contents *[]string, index int) goFunction {

	var goFunc goFunction

	if index >= len(*contents) {
		fmt.Println("Index provided in parseFunction was greater than or equal to the contents length")
		return goFunc
	}

	str := ""

	for i := index; i < len(*contents); i++ {
		str += strings.TrimSpace((*contents)[i])

		if strings.Contains((*contents)[i], "}") {

			re := regexp.MustCompile(`func\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)\s*([^\{]*)\{`) // I wish I understood this regex black magic
			matches := re.FindStringSubmatch(str)

			if len(matches) > 0 {

				goFunc.funcName = matches[1]
				goFunc.funcArgs = strings.Split(matches[2], ",")
				goFunc.funcReturn = matches[3]

				fmt.Printf("Name: %s\nArgs: %s\nReturn %s\n\n\n", goFunc.funcName, goFunc.funcArgs, goFunc.funcReturn)

			}

			return goFunc

		}
	}
	return goFunc
}

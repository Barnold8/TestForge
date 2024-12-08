package main

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

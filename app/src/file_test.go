package main

import (
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"slices"
	"sort"
	"strings"
	"testing"
)

var FILEPERM_test fs.FileMode = 0644
var PATH string = "tests/"

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
	err := os.WriteFile(path, []byte(fileContents), FILEPERM_test)
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
		possibleError := os.Mkdir(path, 0755)
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

				if len(tc.expected) > len(result) {
					for _, value := range tc.expected {

						if slices.Contains(result, value) {
							t.Errorf("%s is in expected\n", value)
						} else {
							t.Errorf("%s is not in expected\n", value)
						}
					}
				} else {
					for _, value := range result {

						if slices.Contains(tc.expected, value) {
							t.Errorf("%s is in expected\n", value)
						} else {
							t.Errorf("%s is not in expected\n", value)
						}
					}
				}
				// t.Errorf("== RESULT ==\n\n%s\n\n == EXPECTED ==\n\n%s\n\n", result, tc.expected)
			}

		})
	}

	// cleanup after testing

	cleanup(PATH)

	// cleanup after testing

}

func TestSeekGoFiles(t *testing.T) {
	// pre testing error catching
	var errors []error
	errors = append(errors, writeTestDirectory("tests"))
	errors = append(errors, writeTestDirectory("tests/test1"))
	errors = append(errors, writeTestDirectory("tests/test1/a"))
	errors = append(errors, writeTestDirectory("tests/test1/a/a"))
	errors = append(errors, writeTestDirectory("tests/test1/a/b"))
	errors = append(errors, writeTestDirectory("tests/test1/a/c"))
	errors = append(errors, writeTestDirectory("tests/test1/b"))
	errors = append(errors, writeTestDirectory("tests/test1/c"))
	errors = append(errors, writeTestDirectory("tests/test1/d"))
	errors = append(errors, writeTestDirectory("tests/test1/e"))
	errors = append(errors, writeTestDirectory("tests/test1/f"))
	errors = append(errors, writeTestDirectory("tests/test1/g"))
	errors = append(errors, writeTestDirectory("tests/test2"))
	errors = append(errors, writeTestDirectory("tests/test3"))
	errors = append(errors, writeTestDirectory("tests/test4"))

	errors = append(errors, writeTestFile(PATH+"main.go", ""))
	errors = append(errors, writeTestFile(PATH+"main.txt", ""))
	errors = append(errors, writeTestFile(PATH+"empty.txt", ""))

	// test1 directory
	errors = append(errors, writeTestFile(PATH+"/test1/go.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/go.go.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/goooooo.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/goooooo_test.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/go.gooooo", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/go.pp", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/go.csv", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/abcsssss.ogg", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/go.helloWorld", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/go.gone", ""))

	// test1/a directory
	errors = append(errors, writeTestFile(PATH+"/test1/a/go.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/a/go.gone", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/a/go.gnome", ""))

	// test1/b directory
	errors = append(errors, writeTestFile(PATH+"/test1/b/go.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/b/g0.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/b/g-o.go", ""))
	// test1/c directory
	errors = append(errors, writeTestFile(PATH+"/test1/c/go.goo", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/c/go_.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/c/go .go", ""))
	// test1/d directory
	errors = append(errors, writeTestFile(PATH+"/test1/d/GO.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/d/gO.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/d/Go.go", ""))

	// test1/e directory
	errors = append(errors, writeTestFile(PATH+"/test1/e/go.goo.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/e/go.g", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/e/go.go.bak", ""))

	// test1/f directory
	errors = append(errors, writeTestFile(PATH+"/test1/f/go.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/f/go_.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/f/go__test.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/f/go .go", ""))

	// test1/g directory
	errors = append(errors, writeTestFile(PATH+"/test1/g/gò.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/g/gó.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/g/g.go", ""))

	// test1/a/a directory
	errors = append(errors, writeTestFile(PATH+"/test1/a/a/go.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/a/a/go_test.go", ""))
	errors = append(errors, writeTestFile(PATH+"/test1/a/a/goNonTest.go", ""))

	// test1/a/b directory
	errors = append(errors, writeTestFile(PATH+"/test1/a/b/go.go", ""))

	// test1/a/c directory
	errors = append(errors, writeTestFile(PATH+"/test1/a/c/go.go", ""))

	for _, err := range errors {

		if err != nil {
			cleanup(PATH)
			t.Errorf("Error - TestReadFileToLines: %s", err)
		}

	}
	// pre testing error catching

	tests := []struct {
		name      string
		path      string
		expected  []string
		overWrite bool
	}{
		{
			"Test Temp directories while igoring existing tests",
			"tests",
			[]string{
				PATH + "main.go",
				PATH + "test1/go.go",
				PATH + "test1/go.go.go",
				PATH + "test1/a/go.go",
				PATH + "test1/b/go.go",
				PATH + "test1/b/g0.go",
				PATH + "test1/b/g-o.go",
				PATH + "test1/c/go_.go",
				PATH + "test1/c/go .go",
				PATH + "test1/d/GO.go",
				PATH + "test1/e/go.goo.go",
				PATH + "test1/f/go.go",
				PATH + "test1/f/go .go",
				PATH + "test1/g/gò.go",
				PATH + "test1/g/gó.go",
				PATH + "test1/g/g.go",
				PATH + "test1/a/a/goNonTest.go",
				PATH + "test1/a/b/go.go",
				PATH + "test1/a/c/go.go0",
			}, false,
		},
		{
			"Test Temp directories while overwriting tests",
			"tests",
			[]string{
				PATH + "main.go",
				PATH + "test1/go.go",
				PATH + "test1/go.go.go",
				PATH + "test1/goooooo.go",
				PATH + "test1/a/go.go",
				PATH + "test1/b/go.go",
				PATH + "test1/b/g0.go",
				PATH + "test1/b/g-o.go",
				PATH + "test1/c/go_.go",
				PATH + "test1/c/go .go",
				PATH + "test1/d/GO.go",
				PATH + "test1/e/go.goo.go",
				PATH + "test1/f/go.go",
				PATH + "test1/f/go_.go",
				PATH + "test1/f/go .go",
				PATH + "test1/g/gò.go",
				PATH + "test1/g/gó.go",
				PATH + "test1/g/g.go",
				PATH + "test1/a/a/go.go",
				PATH + "test1/a/a/goNonTest.go",
				PATH + "test1/a/b/go.go",
				PATH + "test1/a/c/go.go",
			}, true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result, err := seekGoFiles(tc.path, tc.overWrite)

			sort.Strings(result)
			sort.Strings(tc.expected)

			if err != nil {
				cleanup(PATH)
				t.Errorf("Error - TestSeekGoFiles: %s\n", err)
			} else if reflect.DeepEqual(result, tc.expected) != true {
				cleanup(PATH)
				if len(tc.expected) > len(result) {
					for _, value := range tc.expected {

						if slices.Contains(result, value) {
							t.Errorf("%s is in expected\n", value)
						} else {
							t.Errorf("%s is not in expected\n", value)
						}
					}
				} else {
					for _, value := range result {

						if slices.Contains(tc.expected, value) {
							t.Errorf("%s is in expected\n", value)
						} else {
							t.Errorf("%s is not in expected\n", value)
						}
					}
				}
			}
		})
	}

	// cleanup after testing

	cleanup(PATH)

	// cleanup after testing

}

func TestWriteFile(t *testing.T) {
	// pre testing error catching
	var errors []error
	errors = append(errors, writeTestDirectory("tests"))

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
		contents string
		expected []string
	}{
		{
			"main.go write",
			PATH + "main.go",
			"package main\n\nimport \"fmt\"\n\nfunc main(){\n\n\tfmt.Println(\"Hello world\")\n\n}",
			[]string{"package main", "", "import \"fmt\"", "", "func main(){", "", "\tfmt.Println(\"Hello world\")", "", "}"},
		},
		{
			"test.txt write",
			PATH + "test.txt",
			"hello world\ngoodbye universe\nThe quick brown fox jumps over the lazy dog\nThe sun dipped below the horizon, painting the sky in hues of orange and purple.\nA curious cat peeked into the box, only to find it empty.\nThe old bookstore smelled of leather and parchment, a treasure trove of forgotten stories.\nHe poured a steaming cup of coffee and watched the rain tap against the window.\nThe sound of laughter echoed through the park as children chased each other around.\nA soft breeze carried the scent of blooming jasmine through the open window.\nShe flipped through the album, each photo capturing a moment frozen in time.\nThe robot whirred to life, blinking its LED eyes with curiosity.\nMountains loomed in the distance, their peaks kissed by drifting clouds.\nBeneath the waves, a vibrant coral reef thrived with life and color.",
			[]string{"hello world", "goodbye universe", "The quick brown fox jumps over the lazy dog", "The sun dipped below the horizon, painting the sky in hues of orange and purple.", "A curious cat peeked into the box, only to find it empty.", "The old bookstore smelled of leather and parchment, a treasure trove of forgotten stories.", "He poured a steaming cup of coffee and watched the rain tap against the window.", "The sound of laughter echoed through the park as children chased each other around.", "A soft breeze carried the scent of blooming jasmine through the open window.", "She flipped through the album, each photo capturing a moment frozen in time.", "The robot whirred to life, blinking its LED eyes with curiosity.", "Mountains loomed in the distance, their peaks kissed by drifting clouds.", "Beneath the waves, a vibrant coral reef thrived with life and color."},
		},
		{
			"empty.txt write",
			PATH + "empty.txt",
			"",
			[]string{""},
		},
		{
			"test_test.go write",
			PATH + "test_test.go",
			"package main\n\nimport \"fmt\"\n\nfunc TestFunction1(){\n\n}\n\nfunc TestFunction2(){\n\n}\n\nfunc TestFunction3(){\n\n}\n\nfunc TestFunction4(){\n\n}\n\nfunc TestFunction5(){\n\n}\n\n",
			[]string{"package main", "", "import \"fmt\"", "", "func TestFunction1(){", "", "}", "", "func TestFunction2(){", "", "}", "", "func TestFunction3(){", "", "}", "", "func TestFunction4(){", "", "}", "", "func TestFunction5(){", "", "}", "", ""},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := writeFile(tc.path, tc.contents)

			if err != nil {
				cleanup(PATH)
				t.Errorf("Error - TestWriteFile: %s\n", err)
			}
			result, err := readFileToLines(tc.path)

			if err != nil {
				cleanup(PATH)
				t.Errorf("Error - TestWriteFile: %s\n", err)
			}

			if !reflect.DeepEqual(tc.expected, result) {
				cleanup(PATH)
				t.Error("Error - TestWriteFile:\n\n == RESULT ==")
				for _, value := range result {
					cleanup(PATH)
					t.Error(value)
				}
				t.Error("\n\n == EXPECTED ==")
				for _, value := range tc.expected {
					cleanup(PATH)
					t.Error(value)
				}
			}

		})
	}

	cleanup(PATH)

}

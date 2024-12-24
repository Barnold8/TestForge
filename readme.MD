# TestForge

A Go program to generate Go tests!

---

## How to 

### Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/Barnold8/TestForge.git 
   ```
2. Compile the binary
    - Windows 
    ```bat
    go build -o TestForge.exe
    ```
    - Linux
    ```bash
    go build -o TestForge
    ```

### Run the program

To run this program all you need is the binary name and some optional flags


**_Flags:_**

| Name      | Required value      | Usage      |
|-----------------|----------------|----------------|
| --path |  EXAMPLES: <--path D:/Test/Folder> <--path C:/Hello/World>  | Used to find go files in directory and subdirectories|
| --ignore  | EXAMPLES: <--ignore D:/MyProgram/main.go>  <--ignore C:/MyProgram/src/app/man.go C:/MyProgram/src/app/man_test.go > | Used to specify a list of files to ignore in the program |
| --cases | N/A  | Adds basic test cases in given test files |
| --overwrite | N/A  | Ignores any existing test files and overwrites them, alternatively, this can be combined with --ignore to ignore specific files still  |


#### Examples:

1. Running from root directory

```
./TestForge

WARNING: No path was specified. This could lead to unintended results, such as generating tests for directories you did not intend. This is especially risky if you are running this program from the top level of a storage drive.

Try specifying a path with one of the following options:
--path {your path here} or --path={your path here}
Run anyway? y/n:
y
```

This option first warns the user about a possible catastrophe since this program searches all sub directories, read the warning for more info. If the user gives the go ahead, all go files found will have a test file made for them if they dont already have one


2. Specifying directory AND ensuring all files have found have test boilerplate written for them

```bash
./TestForge --path D:/Hello/World --overwrite
```

This will find all possible files within the directories and sub directories of the given directory. It will then write boiler plate test code for ALL of these files found. This means that it will overwrite pre-existing test files, this is due to the --overwrite flag






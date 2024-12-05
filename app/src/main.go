package main

import "fmt"

func main() {

	files, err := seekGoFiles(".")

	if err == nil {

		for _, file := range files {

			fmt.Println(file)

		}

	} else {
		fmt.Printf("Error: %s\n", err)
	}

}

//

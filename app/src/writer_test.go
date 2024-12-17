package main

import("testing")

func TestWriteTests(t *testing.T) {

	tests := []struct {
		name string
		input string
		expected string
	}{
		{"Test 1","input","output"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			// some way to get a result
			result := "this is an example of a result"
			if result != tc.expected{
				 t.Errorf("This is an example of an error!")
			}
		})
	}
}

func TestFunctionArgsToString(t *testing.T) {

	tests := []struct {
		name string
		input string
		expected string
	}{
		{"Test 1","input","output"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			// some way to get a result
			result := "this is an example of a result"
			if result != tc.expected{
				 t.Errorf("This is an example of an error!")
			}
		})
	}
}

func TestGoFunctionsToString(t *testing.T) {

	tests := []struct {
		name string
		input string
		expected string
	}{
		{"Test 1","input","output"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			// some way to get a result
			result := "this is an example of a result"
			if result != tc.expected{
				 t.Errorf("This is an example of an error!")
			}
		})
	}
}

func TestCapitalizeFunctionName(t *testing.T) {

	tests := []struct {
		name string
		input string
		expected string
	}{
		{"Test 1","input","output"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			// some way to get a result
			result := "this is an example of a result"
			if result != tc.expected{
				 t.Errorf("This is an example of an error!")
			}
		})
	}
}

func TestFormatFileName(t *testing.T) {

	tests := []struct {
		name string
		input string
		expected string
	}{
		{"Test 1","input","output"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			// some way to get a result
			result := "this is an example of a result"
			if result != tc.expected{
				 t.Errorf("This is an example of an error!")
			}
		})
	}
}

func TestGatherFiles(t *testing.T) {

	tests := []struct {
		name string
		input string
		expected string
	}{
		{"Test 1","input","output"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			// some way to get a result
			result := "this is an example of a result"
			if result != tc.expected{
				 t.Errorf("This is an example of an error!")
			}
		})
	}
}
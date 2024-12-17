package main

import("testing")

func TestDefaultCLIargs(t *testing.T) {

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

func TestParseArgs(t *testing.T) {

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
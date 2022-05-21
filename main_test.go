package main

import (
	"testing"
)

func TestProg(t *testing.T) {
	type testCase struct {
		name             string
		inputArgs        []string
		expectedOutput   string
		expectedExitCode exitCode
	}

	testCases := []testCase{
		{"no args", nil, USAGE, ExitError},
		{"many args", []string{"there", "are", "many", "args"}, "could not parse uuid there: invalid UUID length: 5", ExitError},
		{"one bad arg", []string{"bad-arg"}, "could not parse uuid bad-arg: invalid UUID length: 7", ExitError},
		{"valid uuid", []string{"6f49a35b-5da9-4e92-bd13-5f7891845e09"}, "b0mjW12pTpK9E194kYReCQ==", ExitSuccess},
		{"valid uuid no dash", []string{"6f49a35b5da94e92bd135f7891845e09"}, "b0mjW12pTpK9E194kYReCQ==", ExitSuccess},
		{"valid uuid v1", []string{"bbda7484-d89a-11ec-9d64-0242ac120002"}, "u9p0hNiaEeydZAJCrBIAAg==", ExitSuccess},
		{"valid uuid v1 no dash", []string{"bbda7484d89a11ec9d640242ac120002"}, "u9p0hNiaEeydZAJCrBIAAg==", ExitSuccess},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			output, code := prog(tt.inputArgs)
			if output != tt.expectedOutput {
				t.Errorf("Expected %q to be equal to %q", output, tt.expectedOutput)
			}
			if code != tt.expectedExitCode {
				t.Errorf("Expected exitCode to be %d", tt.expectedExitCode)
			}
		})
	}
}

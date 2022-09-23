package main

import (
	"bytes"
	"os"
	"testing"
)

func Test_ParseFromStdin(t *testing.T) {
	type testCase struct {
		name             string
		inputArgs        []string
		stdinValue       string
		expectedOutput   string
		expectedExitCode exitCode
	}
	testCases := [...]testCase{
		{"uuid to byte array", nil, "6f49a35b-5da9-4e92-bd13-5f7891845e09", "b0mjW12pTpK9E194kYReCQ==", ExitSuccess},
		{"byte array to uuid", []string{"-r"}, "b0mjW12pTpK9E194kYReCQ==", "6f49a35b-5da9-4e92-bd13-5f7891845e09", ExitSuccess},
		{"empty input and empty stdin", nil, "", USAGE, ExitError},
	}
	buffer := new(bytes.Buffer)
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(buffer.Reset)
			buffer.Write([]byte(tt.stdinValue))
			output, code := prog(tt.inputArgs, buffer)
			if output != tt.expectedOutput {
				t.Errorf("Expected %q to be equal to %q", output, tt.expectedOutput)
			}
			if code != tt.expectedExitCode {
				t.Errorf("Expected exitCode to be %d", tt.expectedExitCode)
			}
		})
	}
}

func Test_ParseUUIDToByteArray(t *testing.T) {
	type testCase struct {
		name             string
		inputArgs        []string
		expectedOutput   string
		expectedExitCode exitCode
	}

	testCases := [...]testCase{
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
			output, code := prog(tt.inputArgs, os.Stdin)
			if output != tt.expectedOutput {
				t.Errorf("Expected %q to be equal to %q", output, tt.expectedOutput)
			}
			if code != tt.expectedExitCode {
				t.Errorf("Expected exitCode to be %d", tt.expectedExitCode)
			}
		})
	}
}

func Test_ParseByteArrayToUUID(t *testing.T) {
	type testCase struct {
		name             string
		inputArgs        []string
		expectedOutput   string
		expectedExitCode exitCode
	}

	testCases := [...]testCase{
		{"no args", nil, USAGE, ExitError},
		{"many args", []string{"-r", "there", "are", "many", "args"}, "could not parse byte string there: illegal base64 data at input byte 4", ExitError},
		{"one bad arg", []string{"-r", "bad-arg"}, "could not parse byte string bad-arg: illegal base64 data at input byte 3", ExitError},
		{"valid base64", []string{"-r", "b0mjW12pTpK9E194kYReCQ=="}, "6f49a35b-5da9-4e92-bd13-5f7891845e09", ExitSuccess},
		{"valid base64 v1", []string{"-r", "u9p0hNiaEeydZAJCrBIAAg=="}, "bbda7484-d89a-11ec-9d64-0242ac120002", ExitSuccess},
		{"invalid base64", []string{"-r", "YmxhaAo="}, "input byte string is not UUID compatible: invalid UUID (got 5 bytes)", ExitError},
		{"invalid base64", []string{"-abcd", ""}, USAGE, ExitError},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			output, code := prog(tt.inputArgs, os.Stdin)
			if output != tt.expectedOutput {
				t.Errorf("Expected %q to be equal to %q", output, tt.expectedOutput)
			}
			if code != tt.expectedExitCode {
				t.Errorf("Expected exitCode to be %d", tt.expectedExitCode)
			}
		})
	}
}

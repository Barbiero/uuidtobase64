package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/uuid"
)

const USAGE = `Outputs a base64-encoded byte array for an UUID

Usage:
	uuidtobase64 [UUID]                          receives an UUID string and outputs a base64-encoded string
	uuidtobase64--reverse [base64_string]        receives a base64-encoded string and outputs an UUID string`

type exitCode uint8

const (
	ExitSuccess exitCode = iota
	ExitError
)

func prog(args []string, stdin io.Reader) (output string, e exitCode) {
	flags := flag.NewFlagSet("uuidtobase64", flag.ContinueOnError)
	reverse := flags.Bool("r", false, "reads a base64-encoded 16 byte array and outputs as an UUID string")
	parseErr := flags.Parse(args)
	if parseErr != nil {
		return USAGE, ExitError
	}

	var input string
	if len(flags.Args()) == 0 {
		in, _ := io.ReadAll(stdin)
		input = strings.TrimSpace(string(in))
	} else {
		input = strings.TrimSpace(flags.Args()[0])
	}
	if len(input) == 0 {
		return USAGE, ExitError
	}

	if *reverse {
		return parseByteArrayToUUID(input)
	}
	return parseUUIDToByteArray(input)

}

func parseByteArrayToUUID(byteString string) (string, exitCode) {
	byteSlice, err := base64.StdEncoding.DecodeString(byteString)
	if err != nil {
		return "could not parse byte string " + byteString + ": " + err.Error(), ExitError
	}

	res, err := uuid.FromBytes(byteSlice)
	if err != nil {
		return "input byte string is not UUID compatible: " + err.Error(), ExitError
	}

	return res.String(), ExitSuccess
}

func parseUUIDToByteArray(uuidStr string) (string, exitCode) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return "could not parse uuid " + uuidStr + ": " + err.Error(), ExitError
	}

	result := base64.StdEncoding.EncodeToString(id[:])
	return result, ExitSuccess
}

func main() {
	output, code := prog(os.Args[1:], os.Stdin)
	out := os.Stdout
	if code != ExitSuccess {
		out = os.Stderr
	}
	_, _ = fmt.Fprintln(out, output)
	os.Exit(int(code))
}

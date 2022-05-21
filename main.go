package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/google/uuid"
)

const USAGE = `Outputs a base64-encoded byte array for an UUID

Usage:
	uuidtobase64 [UUID]`

type exitCode uint8

const (
	ExitSuccess exitCode = iota
	ExitError
)

func prog(args []string) (output string, e exitCode) {
	if len(args) == 0 || len(args[0]) == 0 {
		return USAGE, ExitError
	}

	id, err := uuid.Parse(args[0])
	if err != nil {
		return "could not parse uuid " + args[0] + ": " + err.Error(), ExitError
	}

	result := base64.StdEncoding.EncodeToString(id[:])
	return result, ExitSuccess
}

func main() {
	output, code := prog(os.Args[1:])
	fmt.Println(output)
	os.Exit(int(code))
}

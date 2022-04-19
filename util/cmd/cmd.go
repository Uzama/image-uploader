package cmd

import (
	"errors"
	"imageUploader/domain/globals"
	"os"
)

// read param from command line
func ReadParam() error {
	args := os.Args[1:]

	if len(args) != 1 {
		return errors.New("there should be exatly one argument pass (auth key)")
	}

	// asign to global variable
	globals.AUTH_KEY = args[0]

	return nil
}

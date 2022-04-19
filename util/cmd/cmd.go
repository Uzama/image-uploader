package cmd

import (
	"errors"
	"imageUploader/domain/globals"
	"os"
)

func ReadParam() error {
	args := os.Args[1:]

	if len(args) != 1 {
		return errors.New("there should be exatly one argument")
	}

	globals.AUTH_KEY = args[0]

	return nil
}

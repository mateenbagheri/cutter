package internal

import (
	"fmt"
	"os"

	"github.com/mateenbagheri/cutter/cmd"
)

const (
	errOpenFileFailure = "failed to open file with file address %s"
)

func ProcessCommand(cmd cmd.Command) error {
	// open file
	for _, file := range cmd.Files {
		if err := ProcessFile(file, cmd.Delimiter, cmd.Fields); err != nil {
			return err
		}
	}
	// effect delimiter first
	// effect fields last
	return nil
}

func ProcessFile(fileAddr, delimiter, fields string) error {
	file, err := os.Open(fileAddr)
	if err != nil {
		return fmt.Errorf(errOpenFileFailure, fileAddr)
	}
	defer file.Close()

	return nil
}

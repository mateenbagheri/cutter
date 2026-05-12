package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	scanner := bufio.NewScanner(file)

	// Note: setting buffer size in case there are larger
	// and longer lines in the file
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		err = ProcessLine(line, delimiter, fields)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessLine(line, delimiter, fields string) error {
	if line == "" {
		fmt.Println()
		return nil
	}

	if delimiter == "" {
		delimiter = " "
	}
	words := strings.Split(line, delimiter)

	for _, word := range words {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")
	return nil
}

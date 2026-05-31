package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/mateenbagheri/cutter/cmd"
)

var (
	ErrOpenFileFailure = errors.New("failed to open file with file address")
	ErrFieldStrEmpty   = errors.New("field flag cannot be empty")
)

func ProcessCommand(cmd cmd.Command) error {
	// open file
	for _, file := range cmd.Files {
		if err := ProcessFile(file, cmd.Delimiter, cmd.Field); err != nil {
			return err
		}
	}
	return nil
}

func ProcessFile(fileAddr, delimiter, fieldStr string) error {
	file, err := os.Open(fileAddr)
	if err != nil {
		return fmt.Errorf("failed to open address %s, %w", fileAddr, ErrOpenFileFailure)
	}
	defer file.Close()

	fs, err := NewFieldSelector(fieldStr)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	// Note: setting buffer size in case there are larger
	// and longer lines in the file
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		ProcessLine(line, delimiter, fs)
	}

	return nil
}

func ProcessLine(line, delimiter string, fs *FieldSelector) {
	const separator = " "
	if line == "" {
		fmt.Println()
		return
	}

	if delimiter == "" {
		delimiter = " "
	}

	result, err := fs.Extract(line, delimiter)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}

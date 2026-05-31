package cmd

import (
	"flag"
	"fmt"
)

var (
	ErrFileRequired       = fmt.Errorf("no files passed as argument")
	ErrFieldsFlagRequired = fmt.Errorf("fields flag is required for cutter to function")
)

type Command struct {
	Field     string
	Delimiter string
	Files     []string
}

func (c *Command) ParseFlags(args []string) error {
	fs := flag.NewFlagSet("cutter", flag.ContinueOnError)
	fs.StringVar(&c.Field, "f", "", "fields to return")
	fs.StringVar(&c.Delimiter, "d", " ", "delimiter")
	err := fs.Parse(args)
	if err != nil {
		return err
	}
	c.Files = fs.Args()
	return nil
}

func (c *Command) Validate() error {
	if len(c.Files) == 0 {
		return ErrFileRequired
	}

	if c.Field == "" {
		return ErrFieldsFlagRequired
	}

	return nil
}

func (c *Command) ParseRange() (st, end int, err error) {
	return 0, 0, nil
}

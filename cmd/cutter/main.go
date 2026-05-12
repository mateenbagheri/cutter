package main

import (
	"fmt"
	"os"

	"github.com/mateenbagheri/cutter/cmd"
	"github.com/mateenbagheri/cutter/internal"
)

func main() {
	var command cmd.Command
	if err := command.ParseFlags(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "flag error:", err)
		os.Exit(1)
	}
	if err := command.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, "command validation error:", err)
		os.Exit(2)
	}
	if err := internal.ProcessCommand(command); err != nil {
		fmt.Fprintln(os.Stderr, "process command failed:", err)
		os.Exit(3)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func Run(args []string) int {
	return RunCustom(args, Commands())
}

func RunCustom(args []string, commands map[string]cli.CommandFactory) int {
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	cli := &cli.CLI{
		Args:       args,
		Commands:   commands,
		Version:    Version,
		HelpFunc:   cli.BasicHelpFunc(Name),
		HelpWriter: os.Stdout,
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to execute: %s\n", err.Error())
	}

	return exitCode
}

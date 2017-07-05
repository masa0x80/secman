package main

import (
	"github.com/masa0x80/secman/command"
	"github.com/mitchellh/cli"
)

func Commands() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &command.AddCommand{}, nil
		},
		"delete": func() (cli.Command, error) {
			return &command.DeleteCommand{}, nil
		},
		"list": func() (cli.Command, error) {
			return &command.ListCommand{}, nil
		},
		"sync": func() (cli.Command, error) {
			return &command.SyncCommand{}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Version:  Version,
				Revision: Revision,
				Name:     Name,
			}, nil
		},
	}
}

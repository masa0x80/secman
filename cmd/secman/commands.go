package main

import (
	"github.com/masa0x80/secman/command"
	"github.com/mitchellh/cli"
)

func Commands() map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"save": func() (cli.Command, error) {
			return &command.SaveCommand{}, nil
		},
		"restore": func() (cli.Command, error) {
			return &command.RestoreCommand{}, nil
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

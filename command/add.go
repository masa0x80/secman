package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// AddCommand is the struct for the add command.
type AddCommand struct{}

// Run runs the command.
// The args are the arguments after the command name.
func (c *AddCommand) Run(args []string) int {
	for _, filename := range args {
		localPath, _ := filepath.Abs(filename)
		if !isFile(localPath) {
			fmt.Fprintf(os.Stderr, "SKIP: %s\n", localPath)
			continue
		}
		secretsRoot, err := secretsRoot()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to execute: %s\n", err.Error())
			continue
		}
		remotePath := filepath.Join(secretsRoot, localPath)
		remoteDir := filepath.Dir(remotePath)
		if err := os.MkdirAll(remoteDir, 0777); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to execute: %s\n", err.Error())
		}
		if err := os.Rename(localPath, remotePath); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to execute: %s\n", err.Error())
		}
		if err := os.Symlink(remotePath, localPath); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to execute: %s\n", err.Error())
		} else {
			fmt.Fprintf(os.Stdout, "INFO: Symlink { %s => %s }\n", localPath, remotePath)
		}
	}

	return 0
}

// Synopsis is the short message shown in the 'secman help' output.
func (c *AddCommand) Synopsis() string {
	return "Add symlinks"
}

// Help is the long description shown in the 'secman help add' output.
func (c *AddCommand) Help() string {
	helpText := `
NAME:
		secman add - Add symlinks

USAGE:
		secman add <files>

DESCRIPTION:
		Put target files under secman control.
		Move secret file(s) from /path/to/file to ~/.secrets/path/to/file, and create symlink.
`
	return strings.TrimSpace(helpText)
}

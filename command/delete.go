package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DeleteCommand is the struct for the delete command.
type DeleteCommand struct{}

// Run runs the command.
// The args are the arguments after the command name.
func (c *DeleteCommand) Run(args []string) int {
	for _, filename := range args {
		localPath, _ := filepath.Abs(filename)
		if !isSymlink(localPath) {
			fmt.Fprintf(os.Stderr, "SKIP: %s\n", localPath)
			continue
		}
		secretsRoot, err := secretsRoot()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to execute: %s\n", err.Error())
			continue
		}
		remotePath := filepath.Join(secretsRoot, localPath)
		if !isFile(remotePath) {
			fmt.Fprintf(os.Stderr, "SKIP: %s not found\n", remotePath)
			continue
		}
		if err := os.Remove(localPath); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to execute: %s\n", err.Error())
		}
		if err := os.Rename(remotePath, localPath); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to execute: %s\n", err.Error())
		} else {
			fmt.Fprintf(os.Stdout, "INFO: Move { %s => %s }\n", remotePath, localPath)
		}
	}

	return 0
}

// Synopsis is the short message shown in the 'secman help' output.
func (c *DeleteCommand) Synopsis() string {
	return "Delete symlinks"
}

// Help is the long description shown in the 'secman help version' output.
func (c *DeleteCommand) Help() string {
	helpText := `
NAME:
		secman delete - Delete symlinks

USAGE:
		secman delete [<fileName>...]

DESCRIPTION:
		Make target files out of secman control.
		Remove the symlink, and move secret file(s) from ~/.secrets/path/to/file to /path/to/file.
`
	return strings.TrimSpace(helpText)
}

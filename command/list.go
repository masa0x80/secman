package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ListCommand is the struct for the list command.
type ListCommand struct{}

// Run runs the command.
// The args are the arguments after the command name.
func (c *ListCommand) Run(args []string) int {
	var dirName string

	if len(args) > 0 {
		dirName = args[0]
	} else {
		dirName = "."
	}
	dirPath, _ := filepath.Abs(dirName)
	if !isDir(dirPath) {
		fmt.Fprintf(os.Stderr, "ERROR: Invalid argument\n")
		return 1
	}

	secretsRoot, err := secretsRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to execute: %s\n", err.Error())
		return 1
	}
	remoteDir := filepath.Join(secretsRoot, dirPath)
	for _, path := range traverseFiles(remoteDir) {
		relPath, _ := filepath.Rel(remoteDir, path)
		fmt.Fprintf(os.Stdout, "%s\n", relPath)
	}

	return 0
}

// Synopsis is the short message shown in the 'secman help' output.
func (c *ListCommand) Synopsis() string {
	return "List secret files"
}

// Help is the long description shown in the 'secman help version' output.
func (c *ListCommand) Help() string {
	helpText := `
NAME:
		secman list - list secret files

USAGE:
		secman list [<dirName>]

DESCRIPTION:
		List secret files under secman control.
`
	return strings.TrimSpace(helpText)
}

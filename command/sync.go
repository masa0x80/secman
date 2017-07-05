package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// SyncCommand is the struct for the sync command.
type SyncCommand struct{}

// Run runs the command.
// The args are the arguments after the command name.
func (c *SyncCommand) Run(args []string) int {
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
		localPath := filepath.Join(dirPath, relPath)
		fmt.Println(localPath)

		// Create backup if `relPath` exists
		if _, err := os.Stat(relPath); !os.IsNotExist(err) {
			unixTime := strconv.FormatInt(time.Now().Unix(), 10)
			newPath := localPath + "." + unixTime
			if err := os.Rename(localPath, newPath); err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: Failed to execute: %s\n", err.Error())
			} else {
				fmt.Fprintf(os.Stdout, "INFO: Create backup { %s => %s }\n", relPath, newPath)
			}
		}

		remotePath := filepath.Join(secretsRoot, filepath.Join(dirPath, localPath))
		if err := os.Symlink(remotePath, localPath); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to execute: %s\n", err.Error())
		} else {
			fmt.Fprintf(os.Stdout, "INFO: Symlink { %s => %s }\n", localPath, remotePath)
		}
	}

	return 0
}

// Synopsis is the short message shown in the 'secman help' output.
func (c *SyncCommand) Synopsis() string {
	return "Syncronize secret files recursively"
}

// Help is the long description shown in the 'secman help version' output.
func (c *SyncCommand) Help() string {
	helpText := `
NAME:
		secman sync - syncronize secret files recursively

USAGE:
		secman sync [<dirName>]

DESCRIPTION:
		 Syncronize secret files under secman control recursively.
`
	return strings.TrimSpace(helpText)
}

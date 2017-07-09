package command

import (
	"os"
	"path/filepath"
	"strings"
)

// RestoreCommand is the struct for the restore command.
type RestoreCommand struct{}

// Run runs the command.
// The args are the arguments after the command name.
func (c *RestoreCommand) Run(args []string) int {
	for _, filename := range args {
		localPath, _ := filepath.Abs(filename)
		_, err := os.Lstat(localPath)
		isNotExist := os.IsNotExist(err)
		if !isNotExist && !isSymlink(localPath) {
			Log("Exist", filename)
			continue
		}

		secretsRoot, err := secretsRoot()
		DieIf(err)

		remotePath := filepath.Join(secretsRoot, localPath)
		if !isFile(remotePath) {
			Log("Exist", remotePath)
			continue
		}
		if !isNotExist {
			ErrorIf(os.Remove(localPath))
		}

		if !ErrorIf(os.Rename(remotePath, localPath)) {
			Log("Restore", filename)
		}
	}

	return 0
}

// Synopsis is the short message shown in the 'secman help' output.
func (c *RestoreCommand) Synopsis() string {
	return "Restore symlinks"
}

// Help is the long description shown in the 'secman help restore' output.
func (c *RestoreCommand) Help() string {
	helpText := `
NAME:
		secman restore - Restore files

USAGE:
		secman restore [<fileName>...]

DESCRIPTION:
		Make target files out of secman control.
		Remove the symlink, and move secret file(s) from ~/.secrets/path/to/file to /path/to/file.
`
	return strings.TrimSpace(helpText)
}

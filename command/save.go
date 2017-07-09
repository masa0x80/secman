package command

import (
	"os"
	"path/filepath"
	"strings"
)

// SaveCommand is the struct for the save command.
type SaveCommand struct{}

// Run runs the command.
// The args are the arguments after the command name.
func (c *SaveCommand) Run(args []string) int {
	for _, filename := range args {
		localPath, _ := filepath.Abs(filename)
		if !isFile(localPath) {
			Log("Skip", filename)
			continue
		}

		secretsRoot, err := secretsRoot()
		DieIf(err)

		remotePath := filepath.Join(secretsRoot, localPath)
		remoteDir := filepath.Dir(remotePath)
		ErrorIf(os.MkdirAll(remoteDir, 0777))
		ErrorIf(os.Rename(localPath, remotePath))

		if !ErrorIf(os.Symlink(remotePath, localPath)) {
			Log("Symlink", filename)
		}
	}

	return 0
}

// Synopsis is the short message shown in the 'secman help' output.
func (c *SaveCommand) Synopsis() string {
	return "Create symlinks"
}

// Help is the long description shown in the 'secman help save' output.
func (c *SaveCommand) Help() string {
	helpText := `
NAME:
		secman save - Create symlinks

USAGE:
		secman save [<fileName>...]

DESCRIPTION:
		Put target files under secman control.
		Move secret file(s) from /path/to/file to ~/.secrets/path/to/file, and create symlink.
`
	return strings.TrimSpace(helpText)
}

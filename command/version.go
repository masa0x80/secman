package command

import (
	"bytes"
	"fmt"
	"os"
)

// VersionCommand is the struct for the version command.
type VersionCommand struct {
	Name     string
	Version  string
	Revision string
}

// Run runs the command.
// The args are the arguments after the command name.
func (c *VersionCommand) Run(args []string) int {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "%s version %s", c.Name, c.Version)
	if c.Revision != "" {
		fmt.Fprintf(&versionString, " (%s)", c.Revision)
	}

	fmt.Fprintf(os.Stdout, "%s", versionString.String())
	return 0
}

// Synopsis is the short message shown in the 'secman help' output.
func (c *VersionCommand) Synopsis() string {
	return fmt.Sprintf("Print %s version and quit", c.Name)
}

// Help is the long description shown in the 'secman help version' output.
func (c *VersionCommand) Help() string {
	return "No help topic for `version`"
}

package command

import (
	"os"

	"github.com/motemen/go-colorine"
)

var logger = colorine.NewLogger(
	colorine.Prefixes{
		"Create":  colorine.Notice,
		"Restore": colorine.Notice,
		"Symlink": colorine.Notice,
		"Exist":   colorine.Warn,

		"Error": colorine.Warn,
		"Skip":  colorine.Info,
	},
	colorine.Info,
)

// Log prints messages
func Log(prefix, message string) {
	logger.Log(prefix, message)
}

// ErrorIf prints messages, if error occured
func ErrorIf(err error) bool {
	if err != nil {
		Log("Error", err.Error())
		return true
	}

	return false
}

// DieIf prints messages and exit, if error occured
func DieIf(err error) {
	if err != nil {
		Log("Error", err.Error())
		os.Exit(1)
	}
}

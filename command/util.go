package command

import "os"

func isFile(file string) bool {
	fi, err := os.Lstat(file)
	if os.IsNotExist(err) {
		return false
	}
	isSymlink := fi.Mode()&os.ModeSymlink != 0
	return !isSymlink && fi.Mode().IsRegular()
}

func isSymlink(file string) bool {
	fi, err := os.Lstat(file)
	if os.IsNotExist(err) {
		return false
	}
	return fi.Mode()&os.ModeSymlink != 0
}

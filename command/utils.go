package command

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

func isFile(file string) bool {
	fi, err := os.Lstat(file)
	if os.IsNotExist(err) {
		return false
	}
	isSymlink := fi.Mode()&os.ModeSymlink != 0
	return !isSymlink && fi.Mode().IsRegular()
}

func isDir(file string) bool {
	fi, err := os.Lstat(file)
	if os.IsNotExist(err) {
		return false
	}
	isSymlink := fi.Mode()&os.ModeSymlink != 0
	return !isSymlink && fi.Mode().IsDir()
}

func isSymlink(file string) bool {
	fi, err := os.Lstat(file)
	if os.IsNotExist(err) {
		return false
	}
	return fi.Mode()&os.ModeSymlink != 0
}

func secretsRoot() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	homeDir := u.HomeDir
	return filepath.Join(homeDir, ".secrets"), nil
}

func traverseFiles(dir string) []string {
	var paths []string

	files, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				paths = append(paths, traverseFiles(filepath.Join(dir, file.Name()))...)
				continue
			}
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths
}

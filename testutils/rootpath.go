package testutils

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GetRootPath() (string, error) {
	// Current working directory
	dir, _ := os.Getwd()
	dirSplitted := strings.Split(dir, "/")

	var rootPath string
	var found bool
	// scan for go.mod backwards
	for i := len(dirSplitted); i >= 0; i-- {
		pathToScan := filepath.Join(dirSplitted[:i]...)
		pathToScan = "/" + pathToScan
		err := filepath.Walk(pathToScan, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == "go.mod" {
				found = true
			}
			return nil
		})
		if err != nil {
			return "", err
		}
		if found {
			rootPath = pathToScan
			break
		}
	}
	if !found {
		return "", errors.New("can't find rootpath")
	}
	return rootPath, nil
}

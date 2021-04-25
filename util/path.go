package util

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetProjectRoot get current project's root path
// return path not contain the exec file
func getProjectRoot() string {
	var (
		path string
		err  error
	)
	defer func() {
		if err != nil {
			panic(fmt.Sprintf("GetProjectRoot error :%+v", err))
		}
	}()
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	return path
}

var confDir = "\\config\\config.yaml"

// GetConfPath get configure file path
func GetConfPath() string {
	return getProjectRoot() + confDir
}

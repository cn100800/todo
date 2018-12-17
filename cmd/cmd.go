package cmd

import (
	"os"
)

const (
	FILE_DIR = ".gotask"
)

var (
	workDir  = os.Getenv("HOME") + string(os.PathSeparator) + FILE_DIR
	pendFile = "pending.data"
)

func init() {
	if _, err := os.Stat(workDir); err != nil {
		os.Mkdir(workDir, os.ModePerm)
	}
	os.Chdir(workDir)
	if _, err := os.Stat(pendFile); os.IsNotExist(err) {
		os.Create(pendFile)
	}
}

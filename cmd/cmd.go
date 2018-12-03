package cmd

import (
	"log"
	"os"
)

const (
	FILE_DIR = ".gotask"
)

func MakeDataDir() {

	if err := createFile(); err != nil {
		log.Fatal(err)
	}
}

func createFile() error {
	f := os.Getenv("HOME") + string(os.PathSeparator) + FILE_DIR
	_, err := os.Stat(f)
	if !os.IsNotExist(err) {
		return nil
	}
	if err := os.Mkdir(f, os.ModePerm); err != nil {
		return err
	}
	return nil
}

package cmd

import (
	"os"

	"github.com/satori/go.uuid"
)

func AppendData(s string) error {
	os.Chdir(workDir)
	f, err := os.OpenFile(pendFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return err
	}
	u := uuid.NewV4().String()
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte(s + " " + u + "\n")); err != nil {
		return err
	}
	return nil
}

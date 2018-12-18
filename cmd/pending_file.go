package cmd

import (
	"encoding/json"
	"os"
	"time"

	"github.com/cn100800/todo/task"
	"github.com/satori/go.uuid"
)

func AppendData(p string) error {
	os.Chdir(workDir)
	f, err := os.OpenFile(pendFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return err
	}
	u := uuid.NewV4().String()
	task := task.Task{}
	task.Uuid = u
	task.Project = p
	task.Status = "pending"
	task.Entry = time.Now().Format(time.RFC3339)
	t, err := json.Marshal(task)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(string(t) + "\n"); err != nil {
		return err
	}
	os.Stdout.WriteString("add new task success \n")
	return nil
}

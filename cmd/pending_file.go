package cmd

import (
	"bufio"
	"encoding/json"
	"os"
	"time"

	"github.com/freecracy/todo/task"
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
	l := task.Task{}
	id, err := getId()
	if err != nil {
		return err
	}
	l.Id = id
	l.Uuid = u
	l.Project = p
	l.Status = "pending"
	l.Entry = time.Now().Format(time.RFC3339)
	t, err := json.Marshal(l)
	if err != nil {
		return err
	}
	if _, err := f.WriteString(string(t) + "\n"); err != nil {
		return err
	}
	os.Stdout.WriteString("add new task success \n")
	return nil
}

func getId() (int, error) {
	os.Chdir(workDir)
	f, err := os.Open(pendFile)
	if err != nil {
		return 0, err
	}
	i := 0
	scanner := bufio.NewScanner(f)
	var t task.Task
	for scanner.Scan() {
		i++
		if err := json.Unmarshal([]byte(scanner.Text()), &t); err != nil {
			return 0, err
		}
		if t.Id != i {
			return i, nil
		}
	}
	i++
	return i, nil
}

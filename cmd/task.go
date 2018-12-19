package cmd

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/cn100800/todo/task"
	"github.com/modood/table"
)

func addTask(s string) error {
	if err := AppendData(s); err != nil {
		return err
	}
	return nil
}

func deleteTask(u string) error {
	os.Chdir(workDir)
	f, err := os.OpenFile(pendFile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	var t task.Task
	w := bufio.NewWriter(f)
	i := 0
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &t); err != nil {
			return err
		}
		if t.Uuid == u {
			continue
		}
		i += len(scanner.Bytes())
		w.WriteString(scanner.Text() + "\n")
	}
	f.Truncate(0)
	f.Seek(0, 0)
	w.Flush()
	return nil
}

func listTask() error {
	os.Chdir(workDir)
	f, err := os.Open(pendFile)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	var t task.Task
	var list []task.Task
	for scanner.Scan() {
		if err := json.Unmarshal([]byte(scanner.Text()), &t); err != nil {
			return err
		}
		list = append(list, t)
	}
	table.Output(list)
	return nil
}

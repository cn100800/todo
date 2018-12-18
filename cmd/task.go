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

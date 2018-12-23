package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Gosuri/uitable"
	"github.com/cn100800/todo/task"
)

func addTask(s string) error {
	if err := AppendData(s); err != nil {
		return err
	}
	return nil
}

func allTask() error {
	os.Chdir(workDir)
	f, err := os.Open(doneFile)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	var t task.Task
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("")
	table.AddRow("", "\033[4mID", "\033[0m \033[4mProject", "\033[0m \033[4mUUID\033[0m")
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &t); err != nil {
			return err
		}
		table.AddRow("", t.Id, " "+t.Project, " "+t.Uuid[0:8])
	}
	table.AddRow("")
	fmt.Println(table)
	return nil
}

func deleteTask(u string) error {
	os.Chdir(workDir)
	f, err := os.OpenFile(pendFile, os.O_CREATE|os.O_RDWR, 0644)
	fd, err := os.OpenFile(doneFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(f)
	var t task.Task
	w := bufio.NewWriter(f)
	wd := bufio.NewWriter(fd)
	i := 0
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &t); err != nil {
			return err
		}
		if t.Uuid[0:8] == u {
			wd.WriteString(scanner.Text() + "\n")
			continue
		}
		i += len(scanner.Bytes())
		w.WriteString(scanner.Text() + "\n")
	}
	f.Truncate(0)
	f.Seek(0, 0)
	w.Flush()
	wd.Flush()
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
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("")
	table.AddRow("", "\033[4mID", "\033[0m \033[4mProject", "\033[0m \033[4mUUID\033[0m")
	//var list []task.Task
	for scanner.Scan() {
		if err := json.Unmarshal([]byte(scanner.Text()), &t); err != nil {
			return err
		}
		table.AddRow("", t.Id, " "+t.Project, " "+t.Uuid[0:8])
	}
	table.AddRow("")
	fmt.Println(table)
	return nil
}

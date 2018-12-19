package cmd

import (
	"io"
	"log"
	"os"
)

func init() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			addTask(os.Args[2])
		case "start":
			log.Println("start")
		case "stop":
			log.Println("stop")
		case "delete":
			deleteTask(os.Args[2])
		case "list":
			listTask()
		case "all":
			log.Println("all")
		default:
			help()
		}
	} else {
		help()
	}
}

func help() {
	s := `todo is a gdt tool.

Usage:

	todo <command> [arguments]

The commands are:

	add       Adds a new task
	delete    Deletes the specified task
	list      Most details of tasks
	start     Marks specified task as started
	stop      Removes the 'start' time from a task
	all       All tasks


Use "todo help <command>" for more information about a command.

`
	io.WriteString(os.Stdout, s)
}

package cmd

import (
	"io"
	"log"
	"os"
)

func init() {
	switch os.Args[1] {
	case "add":
		addTask(os.Args[2])
		return
	case "start":
		log.Println("start")
		return
	case "stop":
		log.Println("stop")
		return
	case "delete":
		log.Println("remove")
		return
	case "list":
		log.Println("list")
		return
	case "all":
		log.Println("all")
		return
	case "help":
		help()
		return
	}
}

func help() {
	s := `
todo is a gdt tool.

Usage:

	todo <command> [arguments]

The commands are:

	add       Adds a new task
	start     Marks specified task as started
	stop      Removes the 'start' time from a task
	delete    Deletes the specified task
	list      Most details of tasks
	all       All tasks


Use "todo help <command>" for more information about a command.

`
	io.WriteString(os.Stdout, s)
}

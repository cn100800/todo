package cmd

import (
	"io"
	"log"
	"os"
)

var version string

func main_init() {
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
			allTask()
		case "url":
			switch os.Args[2] {
			case "add":
				addUrl(os.Args[3], os.Args[4])
			case "remove":
				log.Println("url remove")
			case "info":
				log.Println("url info")
			}
		case "version":
			io.WriteString(os.Stdout, version+"\n")
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

工作原因,有想法无实现,轻喷1
https://github.com/freecracy/todo/issues

Version:
	` + version + `

`
	io.WriteString(os.Stdout, s)
}

package task

import (
	"github.com/freecracy/todo/db"
	"github.com/google/uuid"
)

func AddTask(args []string) (bool, error) {
	UUID, _ := uuid.NewUUID()

	var i = map[string]string{
		"id":    UUID.String(),
		"title": args[0],
	}
	db.Insert("todo", i)
	Success("add success!")
	return true, nil
}

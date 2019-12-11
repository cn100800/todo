package task

import (
	"github.com/apex/log"
	"github.com/freecracy/todo/db"
	"github.com/google/uuid"
)

func AddTask(args []string) (bool, error) {
	UUID, _ := uuid.NewUUID()

	var i = map[string]string{
		"short_id": UUID.String()[:8],
		"uuid":     UUID.String(),
		"title":    args[0],
		"status":   "0",
	}
	db.Insert("todo", i)
	log.Info("add success!")
	return true, nil
}

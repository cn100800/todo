package task

import (
	"github.com/apex/log"
	"github.com/freecracy/todo/db"
)

func ReopenTask(id string) {
	db.Update("todo", map[string]string{"status": "0"}, map[string]string{"short_id": id})
	log.Info("reopen success !")
}

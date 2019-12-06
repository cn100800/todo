package task

import "github.com/freecracy/todo/db"

func ReopenTask(id string) {
	db.Update("todo", map[string]string{"status": "0"}, map[string]string{"short_id": id})
	Success("reopen success !")
}
